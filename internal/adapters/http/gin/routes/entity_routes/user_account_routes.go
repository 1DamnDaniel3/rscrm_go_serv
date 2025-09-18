package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/gorm/generic"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/handlers/generic"
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

	r.POST("/user_accounts/register", userHandler.Create)
	r.GET("/user_accounts/:id", userHandler.GetByID)
}
