package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	studentbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_builders"
	studentgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_group_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	studentUCs *studentbuilders.StudentUseCases,
	studentGroupUCs *studentgroupbuilders.StudentGroupUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	handlers := studentbuilders.NewStudentHandlerBuilder(
		studentUCs,
		studentGroupUCs,
	)

	// ================= Routes =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"students",
		authMiddleware,
		tenantMiddleware,
		handlers.CRUDHandler,
	)

	// ===== extra endpoints =====
	protected.GET("/students/search", handlers.SearchStudentsHandler.Search)
	protected.POST("/students/groupedstudents", handlers.GroupedStudentsHandler.GetGroupedStudents)
	protected.POST("/students/createandgroup", handlers.StudentGroupHandler.CreateStudent)

	// ===== nested =====
	protected.GET("/students/:id/clients", handlers.StudentClientsHandler.GetStudentClients)
	protected.GET("/students/:id/groups", handlers.GetGroupsByStudent.GetClientGroups)
	protected.POST("/students/:studId/groups/:groupId", handlers.StudentGroupHandler.CreateRelation)
	protected.DELETE("/students/:studId/groups/:groupId", handlers.StudentGroupHandler.DeleteRelation)
}
