package entityroutes

import (
	sourcepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/source_policies"
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

func SourceRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	sourceRepo := gormentityrepos.NewGormSourceRepository(db)

	// ==/== policies
	crudPolicy := sourcepolicies.NewSourceCrudPolicy()
	sourcePolicy := sourcepolicies.NewSourcePolicies(crudPolicy)

	sourceCrudUC := genericcruduc.NewCRUDUseCase(sourceRepo, sourcePolicy.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.Source,
		dto.SourceCreateUpdateDTO,
		dto.SourceResponseDTO,
	](sourceCrudUC)

	genericrouter.RegisterCRUDRoutes(r, "sources", authMiddleware, tenantMiddleware, genericHandler)
}
