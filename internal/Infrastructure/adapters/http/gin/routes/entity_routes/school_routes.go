package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	schoolbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/school_builders"

	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func SchoolRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *schoolbuilders.SchoolUseCases,
) {

	// ================= Handlers =================
	schoolHandlers := schoolbuilders.NewSchoolHandlerBuilder(
		useCases,
	)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	protected.GET("/schools/:id", schoolHandlers.CRUDHandler.GetByID)
	protected.PATCH("/schools/:id", schoolHandlers.CRUDHandler.Update)
}
