package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	sourcebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/source_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func SourceRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	sourceUCs *sourcebuilders.SourceUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	sourceHandlers := sourcebuilders.NewSourceHandlerBuilder(
		sourceUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"sources",
		authMiddleware,
		tenantMiddleware,
		sourceHandlers.CRUDHandler,
	)
}
