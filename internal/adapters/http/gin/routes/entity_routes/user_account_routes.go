package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/gorm/generic"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/handlers/generic"
	genericRoute "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/routes/entity_routes/generic_handler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserAccountRoutes(r *gin.RouterGroup, db *gorm.DB) {
	userRepo := generic.NewGormRepository[entities.UserAccount](db)
	userHandler := genericHandler.NewGenericHandler[
		entities.UserAccount,
		dto.UserAccountCreateUpdateDTO,
		dto.UserAccountResponseDTO,
	](userRepo)

	genericRoute.RegisterCRUDRoutes(r, "user_accounts", userHandler)
}
