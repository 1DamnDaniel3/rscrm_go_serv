package entityroutes

import (
	userprofilepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_profile_policies"
	profileucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/profile_ucs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	profilehandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user/profile_handlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProfileRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	profileRepo := gormentityrepos.NewGormUserProfileRepo(db)

	// policies
	crudPolicy := userprofilepolicies.NewUserProfieCrudPolicy()
	getSelfProfilePolicy := userprofilepolicies.NewReadProfilePolicy()

	profilePolicy := userprofilepolicies.NewUserProfilePolicy(
		crudPolicy,
		getSelfProfilePolicy,
	)

	// crud uc
	// profileCrudUC := genericcruduc.NewCRUDUseCase(profileRepo, profilePolicy.CRUD)

	// ==/== get self profile
	getSelfProfileUC := profileucs.NewGetSelfProfileUC(profileRepo, *profilePolicy)
	getSelfProfileHandler := profilehandlers.NewGetSelfProfileHandler(getSelfProfileUC)

	// genericHandler := generichandler.NewGenericHandler[
	// 	entities.UserProfile,
	// 	dto.UserProfileCreateUpdateDTO,
	// 	dto.UserProfileResponseDTO,
	// ](profileCrudUC)

	// genericrouter.RegisterCRUDRoutes(r, "user_profiles", authMiddleware, tenantMiddleware, genericHandler)

	protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)

	protected.GET("/user_account/:id/profile", getSelfProfileHandler.GetSelfProfile)

}
