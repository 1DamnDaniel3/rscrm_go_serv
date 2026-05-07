package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	statusbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/status_builders"

	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func StatusRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	statusUCs *statusbuilders.StatusUseCases,
) {

	// ================= Handlers =================
	statusHandlers := statusbuilders.NewStatusHandlerBuilder(
		statusUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"statuses",
		app.AuthMiddleware,
		app.TenantMiddleware,
		statusHandlers.CRUDHandler,
	)
}
