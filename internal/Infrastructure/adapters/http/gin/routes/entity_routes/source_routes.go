package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	sourcebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/source_builders"

	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func SourceRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	sourceUCs *sourcebuilders.SourceUseCases,
) {

	// ================= Handlers =================
	sourceHandlers := sourcebuilders.NewSourceHandlerBuilder(
		sourceUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"sources",
		app.AuthMiddleware,
		app.TenantMiddleware,
		sourceHandlers.CRUDHandler,
	)
}
