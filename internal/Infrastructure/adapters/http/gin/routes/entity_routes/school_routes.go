package entityroutes

import (
	schoolpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/school_policies"
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

func SchoolRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	schoolRepo := gormentityrepos.NewGormSchoolRepo(db)
	// ==/== policies
	schoolCrudPolicy := schoolpolicies.NewSchoolCrudPolicy()
	schoolPolicies := schoolpolicies.NewSchoolPolicies(schoolCrudPolicy)

	schoolCrudUC := genericcruduc.NewCRUDUseCase(schoolRepo, schoolPolicies.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.School,
		dto.SchoolCreateUpdateDTO,
		dto.SchoolResponseDTO,
	](schoolCrudUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "schools", authMiddleware, tenantMiddleware, genericHandler)
	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)

	protected.GET("/schools/:id", genericHandler.GetByID)
}
