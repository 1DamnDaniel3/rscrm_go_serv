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
	// REPO
	students_repo := gormentityrepos.NewGormStudentsRepo(db)
	studentsQueryService := gormentityrepos.NewGormStudentQueryService(db)
	studentGroups_repo := gormentityrepos.NewGormStudentGroupsRepo(db)

	// generic
	genericStudentsHandler := genericHandler.NewGenericHandler[ // student
		entities.Student,
		dto.StudentCreateUpdateDTO,
		dto.StudentResponseDTO,
	](students_repo)

	// UCs and Handlers

	// === student-group ===
	createAndGroupUC := studentgroupUCs.NewCreateStudentUC(tx, students_repo, studentGroups_repo) // create
	relationCRUDucs := studentgroupUCs.NewStudentGroupCRUDucs(studentGroups_repo)                 // crud

	getGroupedUC := studentucs.NewGroupedStudentsUC(students_repo) // getGrouped
	getGroupedHandler := studenthandlers.NewGetGroupedStudentsHandler(getGroupedUC)

	getGroupsByStudentUC := studentucs.NewGetStudentGroupUC(studentsQueryService) // get groups by student
	getGroupsByStudentHandler := studenthandlers.NewGetStudentGroupsHandler(getGroupsByStudentUC)

	searchStudentsUC := studentucs.NewSearchStudentsUC(studentsQueryService)
	searchHandler := studenthandlers.NewSearchStudentHandler(searchStudentsUC)

	studentGroupHandler := studentgroupHandlers.NewStudentGroupHandler(
		createAndGroupUC, relationCRUDucs) // student-group handler

	// students/{id}/clients
	studnetClientsUC := studentclientUCs.NewGetStudentClientsUC(studentsQueryService)
	studentClientsHandler := studentclientHandlers.NewStudentClientsHandler(studnetClientsUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "students", authMiddleware, tenantMiddleware, genericStudentsHandler)

	protected.GET("/students/search", searchHandler.Search)
	protected.POST("/students/groupedstudents", getGroupedHandler.GetGroupedStudents)
	protected.POST("/students/createandgroup", studentGroupHandler.CreateStudent)

	// ==== nested routes

	// student_clients
	protected.GET("/students/:id/clients", studentClientsHandler.GetStudentClients)
	// student_groups
	protected.GET("/students/:id/groups", getGroupsByStudentHandler.GetClientGroups)
	protected.POST("/students/:studId/groups/:groupId", studentGroupHandler.CreateRelation)
	protected.DELETE("/students/:studId/groups/:groupId", studentGroupHandler.DeleteRelation)
}
