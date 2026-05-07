package leadgroupbuilders

import (
	leadgroupucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs/leadGroupUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type LeadGroupsUseCaseBuilder struct {
	CRUD       genericcruduc.ICRUDUseCase[entities.LeadGroup]
	CreateLead leadgroupucs.CreateLeadUC
}

func NewLeadGroupsUseCaseBuilder(
	tx services.Transaction,
	leadModule *modules.LeadModule,
	leadGroupModule *modules.LeadGroupsModule,
) *LeadGroupsUseCaseBuilder {

	return &LeadGroupsUseCaseBuilder{
		// CRUD

		CRUD: genericcruduc.NewCRUDUseCase(
			leadGroupModule.LeadGroupsRepo,
			leadGroupModule.LeadGroupPolicies.CRUD,
		),

		CreateLead: *leadgroupucs.NewCreateLeadUC(
			tx,
			leadModule.LeadRepo,
			leadGroupModule.LeadGroupsRepo,
			leadModule.LeadPolicies.CRUD,
			leadGroupModule.LeadGroupPolicies.CRUD,
		),
	}
}
