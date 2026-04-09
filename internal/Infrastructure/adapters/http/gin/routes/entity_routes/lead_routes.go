package entityroutes

import (
	leadpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lead_policies"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs"
	leadgroupucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs/leadGroupUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	leadgroupHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers/leadGroupHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"

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
	// ==/== repo
	leadRepo := gormentityrepos.NewGormLeadsRepo(db)
	leadGroupsRepo := gormentityrepos.NewGormLeadGroupsRepo(db)

	// ==/== policies
	crudPolicy := leadpolicies.NewLeadCrudPolicy()
	leadPolicies := leadpolicies.NewLeadPolicies(crudPolicy)

	crudUC := genericcruduc.NewCRUDUseCase(leadRepo, leadPolicies.CRUD)

	groupedLeadsUC := leadUCs.NewGroupedLeadsUC(leadRepo) // Grouped Leads
	getGroupedLeadsHandler := leadhandlers.NewGetGroupedLeadsHandler(groupedLeadsUC)

	// lead-groups
	createLeadsUC := leadgroupucs.NewCreateLeadUC(tx, leadRepo, leadGroupsRepo) // CreateLead
	leadGroupCrudUC := leadgroupucs.NewLeadGroupCRUDucs(leadGroupsRepo)
	leadGroupCrudHandler := leadgroupHandlers.NewLeadGroupHandler(createLeadsUC, leadGroupCrudUC)

	genericHandler := generichandler.NewGenericHandler[
		entities.Lead,
		dto.LeadCreateUpdateDTO,
		dto.LeadResponseDTO,
	](crudUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "leads", authMiddleware, tenantMiddleware, genericHandler)
	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)

	protected.POST("/leads/groupedleads", getGroupedLeadsHandler.GetGroupedLeads)
	protected.POST("/leads/createandgroup", leadGroupCrudHandler.CreateLead)

	// nested
	protected.POST("/leads/:leadId/groups/:groupId", leadGroupCrudHandler.CreateRelation)
	protected.DELETE("/leads/:leadId/groups/:groupId", leadGroupCrudHandler.DeleteRelation)

}
