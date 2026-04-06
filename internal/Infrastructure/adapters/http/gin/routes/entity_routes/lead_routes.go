package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs"
	leadgroupucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs/leadGroupUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	leadgroupHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers/leadGroupHandlers"

	leadhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	leadRepo := gormentityrepos.NewGormLeadsRepo(db)
	leadGroupsRepo := gormentityrepos.NewGormLeadGroupsRepo(db)

	groupedLeadsUC := leadUCs.NewGroupedLeadsUC(leadRepo) // Grouped Leads
	getGroupedLeadsHandler := leadhandlers.NewGetGroupedLeadsHandler(groupedLeadsUC)

	// lead-groups
	createLeadsUC := leadgroupucs.NewCreateLeadUC(tx, leadRepo, leadGroupsRepo) // CreateLead
	leadGroupCrudUC := leadgroupucs.NewLeadGroupCRUDucs(leadGroupsRepo)
	leadGroupCrudHandler := leadgroupHandlers.NewLeadGroupHandler(createLeadsUC, leadGroupCrudUC)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Lead,
	// 	dto.LeadCreateUpdateDTO,
	// 	dto.LeadResponseDTO,
	// ](leadCrudUC)

	// protected := genericrouter.RegisterCRUDRoutes(r, "leads", authMiddleware, tenantMiddleware, genericHandler)
	protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)
	protected.POST("/leads/groupedleads", getGroupedLeadsHandler.GetGroupedLeads)
	protected.POST("/leads/createandgroup", leadGroupCrudHandler.CreateLead)

	// nested
	protected.POST("/leads/:leadId/groups/:groupId", leadGroupCrudHandler.CreateRelation)
	protected.DELETE("/leads/:leadId/groups/:groupId", leadGroupCrudHandler.DeleteRelation)

}
