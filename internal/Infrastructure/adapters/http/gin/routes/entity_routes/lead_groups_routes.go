package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	leadgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_group_builders"

	"github.com/gin-gonic/gin"
)

func LeadGroupsRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	leadGroupUCs *leadgroupbuilders.LeadGroupsUseCaseBuilder,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ================= Handlers =================
	leadGroupHandlers := leadgroupbuilders.NewLeadGroupsHandlerBuilder(
		leadGroupUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"lead_groups",
		authMiddleware,
		tenantMiddleware,
		leadGroupHandlers.CRUDHandler,
	)
}
