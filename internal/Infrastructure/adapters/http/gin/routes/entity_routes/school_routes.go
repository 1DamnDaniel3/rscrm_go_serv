package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SchoolRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	schoolRepo := gormentityrepos.NewGormSchoolRepo(db)

	genericHandler := genericHandler.NewGenericHandler[
		entities.School,
		dto.SchoolCreateUpdateDTO,
		dto.SchoolResponseDTO,
	](schoolRepo)

	genericrouter.RegisterCRUDRoutes(r, "schools", genericHandler)
}
