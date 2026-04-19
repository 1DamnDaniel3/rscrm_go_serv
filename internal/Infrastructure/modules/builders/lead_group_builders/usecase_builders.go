package leadgroupbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type LeadGroupsUseCaseBuilder struct {
	CRUD genericcruduc.ICRUDUseCase[entities.LeadGroup]
}

func NewLeadGroupsUseCaseBuilder(
	leadGroupModule *modules.LeadGroupsModule,
) *LeadGroupsUseCaseBuilder {

	return &LeadGroupsUseCaseBuilder{
		// CRUD

		CRUD: genericcruduc.NewCRUDUseCase(
			leadGroupModule.LeadGroupsRepo,
			leadGroupModule.LeadGroupPolicies.CRUD,
		),
	}
}
