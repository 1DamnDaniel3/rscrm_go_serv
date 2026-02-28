package entityroutes

import (
	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
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
	genericStudentsHandler := genericHandler.NewGenericHandler[
		entities.Student,
		dto.StudentCreateUpdateDTO,
		dto.StudentResponseDTO,
	](students_repo)
	//groupedstudents
	getGroupedUC := studentucs.NewGroupedStudentsUC(students_repo)
	groupedHandler := studenthandlers.NewGetGroupedHandler(getGroupedUC)

	// createandgroup
	createAndGroupUC := studentucs.NewCreateStudentUC(tx, students_repo, studentGroups_repo)
	createAndGroupHandler := studenthandlers.NewCreateStudentHandler(createAndGroupUC)

	// students/{id}/clients
	studnetClientsUC := studentucs.NewGetStudentClientsUC(gormStudentsQueryService)
	studentClientsHandler := studenthandlers.NewStudentClientsHandler(studnetClientsUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "students", authMiddleware, tenantMiddleware, genericStudentsHandler)
	protected.GET("/students/:id/clients", studentClientsHandler.GetStudentClients)
	protected.POST("/students/groupedstudents", groupedHandler.GetGroupedStudents)
	protected.POST("/students/createandgroup", createAndGroupHandler.CreateStudent)

	// nested routes

}
