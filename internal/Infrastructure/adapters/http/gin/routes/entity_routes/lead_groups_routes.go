package entityroutes

import (
	leadgrouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lead_group_policies"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadGroupsRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	leadGroupsRepo := gormentityrepos.NewGormLeadGroupsRepo(db)

	// ==/== policies

	crudPolicy := leadgrouppolicies.NewLeadGroupCrudPolicy()
	leadGroupPolicy := leadgrouppolicies.NewLeadGroupPolicies(crudPolicy)

	crudUC := genericcruduc.NewCRUDUseCase(leadGroupsRepo, leadGroupPolicy.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.LeadGroup,
		dto.LeadGroupCreateUpdateDTO,
		dto.LeadGroupResponseDTO,
	](crudUC)

	genericrouter.RegisterCRUDRoutes(r, "lead_groups", authMiddleware, tenantMiddleware, genericHandler)
}
