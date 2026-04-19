package leadgroupbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type LeadGroupsHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.LeadGroup,
		dto.LeadGroupCreateUpdateDTO,
		dto.LeadGroupResponseDTO,
	]
}

func NewLeadGroupsHandlerBuilder(useCases *LeadGroupsUseCaseBuilder) *LeadGroupsHandlerBuilder {
	return &LeadGroupsHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.LeadGroup,
			dto.LeadGroupCreateUpdateDTO,
			dto.LeadGroupResponseDTO,
		](useCases.CRUD),
	}
}
