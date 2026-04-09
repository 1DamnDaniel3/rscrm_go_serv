package entityroutes

import (
	statuspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/status_policies"
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

func StatusRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	statusRepo := gormentityrepos.NewGormStatusRepository(db)

	// ==/== repo
	crudPolicy := statuspolicies.NewStatusCrudPolicy()
	statusPolicies := statuspolicies.NewStatusPolicies(crudPolicy)

	statusCrudUC := genericcruduc.NewCRUDUseCase(statusRepo, statusPolicies.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.Status,
		dto.StatusCreateUpdateDTO,
		dto.StatusResponseDTO,
	](statusCrudUC)

	genericrouter.RegisterCRUDRoutes(r, "statuses", authMiddleware, tenantMiddleware, genericHandler)
}
