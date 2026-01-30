package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadGroupsRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
) {
	lead_groups_repo := gormentityrepos.NewGormLeadGroupsRepo(db)

	genericHandler := generic.NewGenericHandler[
		entities.LeadGroup,
		dto.LeadGroupCreateUpdateDTO,
		dto.LeadGroupResponseDTO,
	](lead_groups_repo)

	genericrouter.RegisterCRUDRoutes(r, "lead_groups", authMiddleware, genericHandler)
}
