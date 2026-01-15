package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserProfileRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	profileRepo := gormentityrepos.NewGormUserProfileRepo(db)

	genericHandler := generic.NewGenericHandler[
		entities.UserProfile,
		dto.UserProfileCreateUpdateDTO,
		dto.UserProfileResponseDTO,
	](profileRepo)

	genericrouter.RegisterCRUDRoutes(r, "user_profiles", genericHandler)
}
