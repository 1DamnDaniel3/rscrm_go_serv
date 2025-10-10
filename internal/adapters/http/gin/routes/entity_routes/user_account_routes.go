package entityroutes

import (
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/handlers/generic"
	genericRoute "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/routes/entity_routes/generic_handler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports"
	"github.com/gin-gonic/gin"
)

func UserAccountRoutes(r *gin.RouterGroup, userRepo ports.UserAccountRepository) {
	userHandler := genericHandler.NewGenericHandler[
		entities.UserAccount,
		dto.UserAccountCreateUpdateDTO,
		dto.UserAccountResponseDTO,
	](userRepo)

	genericRoute.RegisterCRUDRoutes(r, "user_accounts", userHandler)
	// r.POST("/user_accounts/create")
}
