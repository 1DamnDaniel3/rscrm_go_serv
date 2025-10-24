package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	statusRepo := adapters.NewGormStatusRepository(db)

	genericHandler := generic.NewGenericHandler[
		entities.Status,
		dto.StatusCreateUpdateDTO,
		dto.StatusResponseDTO,
	](statusRepo)

	genericrouter.RegisterCRUDRoutes(r, "statuses", genericHandler)
}
