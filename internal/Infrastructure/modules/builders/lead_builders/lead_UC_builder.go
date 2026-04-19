package leadbuilders

import (
	leaducs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs"
	leadgroupucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/leadUCs/leadGroupUCs"

	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type LeadUseCases struct {
	CRUD                 genericcruduc.ICRUDUseCase[entities.Lead]
	GroupedLeads         leaducs.IGroupedLeadsUC
	CreateLead           leadgroupucs.ICreateLeadUC
	LeadGroupCRUD        leadgroupucs.ILeadGroupCRUDucs
	LeadGroupGenericCRUD genericcruduc.ICRUDUseCase[entities.LeadGroup]
}

func NewLeadUseCasesBuilder(
	tx services.Transaction,
	leadModule *modules.LeadModule,
	leadGroupModule *modules.LeadGroupsModule,
) *LeadUseCases {

	return &LeadUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			leadModule.LeadRepo,
			leadModule.LeadPolicies.CRUD,
		),

		// ================= GROUPED LEADS =================
		GroupedLeads: leaducs.NewGroupedLeadsUC(
			leadModule.LeadRepo,
		),

		// ================= CREATE LEAD =================
		CreateLead: leadgroupucs.NewCreateLeadUC(
			tx,
			leadModule.LeadRepo,
			leadGroupModule.LeadGroupsRepo,
			leadModule.LeadPolicies.CRUD,
			leadGroupModule.LeadGroupPolicies.CRUD,
		),

		// ================= LEAD-GROUP CRUD =================
		LeadGroupCRUD: leadgroupucs.NewLeadGroupCRUDucs(
			leadGroupModule.LeadGroupsRepo,
			leadGroupModule.LeadGroupPolicies.CRUD,
		),
		LeadGroupGenericCRUD: genericcruduc.NewCRUDUseCase(
			leadGroupModule.LeadGroupsRepo,
			leadModule.LeadPolicies.CRUD,
		),
	}
}
