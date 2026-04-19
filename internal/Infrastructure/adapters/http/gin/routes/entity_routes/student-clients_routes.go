package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	studentclientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_clients_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func StudentClientRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	studCliUCs *studentclientbuilders.StudentClientUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	handlers := studentclientbuilders.NewStudentClientHandlerBuilder(
		studCliUCs,
	)

	// ================= Routes =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"student-clients",
		authMiddleware,
		tenantMiddleware,
		handlers.CRUDHandler,
	)

	protected.POST("/student-clients/createandgetBO", handlers.CreateRelationHandler.CreateRel)
}
