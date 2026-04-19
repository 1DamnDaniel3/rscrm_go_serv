package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	schoolbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/school_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func SchoolRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *schoolbuilders.SchoolUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	schoolHandlers := schoolbuilders.NewSchoolHandlerBuilder(
		useCases,
	)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		authMiddleware,
		tenantMiddleware,
	)

	protected.GET("/schools/:id", schoolHandlers.CRUDHandler.GetByID)
}
