package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	groupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/group_builders"

	"github.com/gin-gonic/gin"
)

func GroupRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	groupUCs *groupbuilders.GroupUseCases,
) {
	// ================= Handlers =================
	groupHandlers := groupbuilders.NewGroupHandlerBuilder(
		groupUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"groups",
		app.AuthMiddleware,
		app.TenantMiddleware,
		groupHandlers.CRUDHandler,
	)
}
