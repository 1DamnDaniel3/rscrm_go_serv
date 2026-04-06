package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProfileRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// profileRepo := gormentityrepos.NewGormUserProfileRepo(db)

	// policies
	// crudPolicy := userprofilepolicies.NewUserProfieCrudPolicy()
	// profilePolicy := userprofilepolicies.NewUserProfilePolicy(crudPolicy)
	// profileCrudUC := genericcruduc.NewCRUDUseCase(profileRepo, profilePolicy.CRUD)

	// genericHandler := generichandler.NewGenericHandler[
	// 	entities.UserProfile,
	// 	dto.UserProfileCreateUpdateDTO,
	// 	dto.UserProfileResponseDTO,
	// ](profileCrudUC)

	// genericrouter.RegisterCRUDRoutes(r, "user_profiles", authMiddleware, tenantMiddleware, genericHandler)

}
