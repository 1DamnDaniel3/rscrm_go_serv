package leadbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	leadhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers"
	leadgrouphandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/leadHandlers/leadGroupHandlers"
	leadgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_group_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type LeadHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Lead,
		dto.LeadCreateUpdateDTO,
		dto.LeadResponseDTO,
	]
	GroupedLeadsHandler *leadhandlers.GetGroupedLeadsHandler
	LeadGroupHandler    *leadgrouphandlers.LeadGroupHandler
}

func NewLeadHandlerBuilder(
	leadUCs *LeadUseCases,
	leadGroupUCs *leadgroupbuilders.LeadGroupsUseCaseBuilder,
) *LeadHandlerBuilder {
	return &LeadHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Lead,
			dto.LeadCreateUpdateDTO,
			dto.LeadResponseDTO,
		](leadUCs.CRUD),

		// ================= GROUPED =================
		GroupedLeadsHandler: leadhandlers.NewGetGroupedLeadsHandler(
			leadUCs.GroupedLeads,
		),

		// ================= LEAD-GROUP =================
		LeadGroupHandler: leadgrouphandlers.NewLeadGroupHandler(
			leadUCs.CreateLead,
			leadUCs.LeadGroupCRUD,
			leadGroupUCs.CRUD,
		),
	}
}
