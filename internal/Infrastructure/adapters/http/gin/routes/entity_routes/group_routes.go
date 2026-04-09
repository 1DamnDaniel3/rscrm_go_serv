package entityroutes

import (
	grouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/group_policies"
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

func GroupRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	groupRepo := gormentityrepos.NewGormGroupRepository(db)

	// ==/== policies
	crudPolicy := grouppolicies.NewGroupCrudPolicy()

	groupPolicies := grouppolicies.NewGroupPolicies(crudPolicy)

	groupCrudUC := genericcruduc.NewCRUDUseCase(groupRepo, groupPolicies.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.Group,
		dto.GroupCreateUpdateDTO,
		dto.GroupResponseDTO,
	](groupCrudUC)

	genericrouter.RegisterCRUDRoutes(r, "groups", authMiddleware, tenantMiddleware, genericHandler)
}
