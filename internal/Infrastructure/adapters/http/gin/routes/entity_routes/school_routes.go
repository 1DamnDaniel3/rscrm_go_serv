package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/school"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SchoolRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	schoolRepo := adapters.NewGormSchoolRepo(db)

	genericHandler := genericHandler.NewGenericHandler[
		entities.School,
		dto.SchoolCreateUpdateDTO,
		dto.SchoolResponseDTO,
	](schoolRepo)
	getByIDHandler := school.NewSchoolHandler(schoolRepo)

	genericrouter.RegisterCRUDRoutes(r, "schools", genericHandler)
	r.GET("/schools/getoneschool/:id", getByIDHandler.GetByID)
}
