package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	leadhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
) {
	leadRepo := gormentityrepos.NewGormLeadsRepo(db)
	leadGroupsRepo := gormentityrepos.NewGormLeadGroupsRepo(db)

	// Grouped Leads
	groupedLeadsUC := leadUCs.NewGroupedLeadsUC(leadRepo)
	groupedLeadsHandler := leadhandlers.NewGroupedLeadsHandler(groupedLeadsUC)

	// CreateLead
	createLeadsUC := leadUCs.NewCreateLeadUC(tx, leadRepo, leadGroupsRepo)
	createLeadHandler := leadhandlers.NewCreateLeadHandler(createLeadsUC)

	genericHandler := generic.NewGenericHandler[
		entities.Lead,
		dto.LeadCreateUpdateDTO,
		dto.LeadResponseDTO,
	](leadRepo)

	genericrouter.RegisterCRUDRoutes(r, "leads", authMiddleware, genericHandler)
	r.POST("leads/groupedleads", groupedLeadsHandler.GetGroupedLeads)
	r.POST("leads/createandgroup", createLeadHandler.CreateLead)

}
