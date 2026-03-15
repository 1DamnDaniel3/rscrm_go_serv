package entityroutes

import (
	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	studentclientUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs/student_clientUCs"
	studentgroupUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs/student_groupUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	studentclientHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers/studentClientHandlers.go"
	studentgroupHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers/studentGroupHandlers.go"

	studenthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StudentRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	students_repo := gormentityrepos.NewGormStudentsRepo(db)
	gormStudentsQueryService := gormentityrepos.NewGormStudentQueryService(db)
	studentGroups_repo := gormentityrepos.NewGormStudentGroupsRepo(db)
	// generic
	genericStudentsHandler := genericHandler.NewGenericHandler[ // student
		entities.Student,
		dto.StudentCreateUpdateDTO,
		dto.StudentResponseDTO,
	](students_repo)

	// === student-group === studentucs
	createAndGroupUC := studentgroupUCs.NewCreateStudentUC(tx, students_repo, studentGroups_repo) // create

	getGroupedUC := studentucs.NewGroupedStudentsUC(students_repo) // getGrouped
	getGroupedHandler := studenthandlers.NewGetGroupedStudentsHandler(getGroupedUC)

	createStudGroupRelationUC := studentgroupUCs.NewCreateStudentGroupRelationUC(studentGroups_repo)

	studentGroupHandler := studentgroupHandlers.NewStudentGroupHandler(
		createAndGroupUC, createStudGroupRelationUC) // student-group handler

	// students/{id}/clients
	studnetClientsUC := studentclientUCs.NewGetStudentClientsUC(gormStudentsQueryService)
	studentClientsHandler := studentclientHandlers.NewStudentClientsHandler(studnetClientsUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "students", authMiddleware, tenantMiddleware, genericStudentsHandler)
	protected.POST("/students/groupedstudents", getGroupedHandler.GetGroupedStudents)
	protected.POST("/students/createandgroup", studentGroupHandler.CreateStudent)

	// nested routes
	protected.GET("/students/:id/clients", studentClientsHandler.GetStudentClients)
	protected.POST("/students/:studId/groups/:groupId", studentGroupHandler.CreateRelation)
}
