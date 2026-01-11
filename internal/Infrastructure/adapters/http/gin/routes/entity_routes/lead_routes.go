package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	leadhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
) {
	leadRepo := adapters.NewGormLeadsRepo(db)

	// Grouped Leads
	groupedLeadsUC := leadUCs.NewGroupedLeadsUC(leadRepo)
	groupedLeadsHandler := leadhandlers.NewGroupedLeadsHandler(groupedLeadsUC)

	genericHandler := generic.NewGenericHandler[
		entities.Lead,
		dto.LeadCreateUpdateDTO,
		dto.LeadResponseDTO,
	](leadRepo)

	genericrouter.RegisterCRUDRoutes(r, "leads", genericHandler)
	r.POST("leads/groupedleads", groupedLeadsHandler.GetGroupedLeads)

}
