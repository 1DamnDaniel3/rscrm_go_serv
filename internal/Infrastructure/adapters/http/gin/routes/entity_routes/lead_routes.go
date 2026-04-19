package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	leadbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_builders"
	leadgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_group_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func LeadRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	leadUCs *leadbuilders.LeadUseCases,
	leadGroupUCs *leadgroupbuilders.LeadGroupsUseCaseBuilder,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	handlers := leadbuilders.NewLeadHandlerBuilder(
		leadUCs,
		leadGroupUCs,
	)

	// ================= Routes =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"leads",
		authMiddleware,
		tenantMiddleware,
		handlers.CRUDHandler,
	)

	protected.POST("/leads/groupedleads", handlers.GroupedLeadsHandler.GetGroupedLeads)
	protected.POST("/leads/createandgroup", handlers.LeadGroupHandler.CreateLead)

	// nested
	protected.POST("/leads/:leadId/groups/:groupId", handlers.LeadGroupHandler.CreateRelation)
	protected.DELETE("/leads/:leadId/groups/:groupId", handlers.LeadGroupHandler.DeleteRelation)
}
