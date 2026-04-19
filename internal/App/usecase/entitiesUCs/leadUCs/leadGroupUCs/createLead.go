package leadgroupucs

import (
	"context"

	leadpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lead_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateLeadUC struct {
	tx             services.Transaction
	leadRepo       entitiesrepos.LeadsRepository
	leadGroupsRepo entitiesrepos.LeadGroupsRepo

	leadPolicy      leadpolicies.ILeadCrudPolicy
	leadGroupPolicy leadpolicies.ILeadCrudPolicy
}

type ICreateLeadUC interface {
	Execute(ctx context.Context, lead *entities.Lead, leadGroup *entities.LeadGroup) error
}

func NewCreateLeadUC(
	tx services.Transaction,
	leadRepo entitiesrepos.LeadsRepository,
	leadGroupsRepo entitiesrepos.LeadGroupsRepo,
	leadPolicy leadpolicies.ILeadCrudPolicy,
	leadGroupPolicy leadpolicies.ILeadCrudPolicy,
) *CreateLeadUC {
	return &CreateLeadUC{
		tx:             tx,
		leadRepo:       leadRepo,
		leadGroupsRepo: leadGroupsRepo,
	}
}

func (uc *CreateLeadUC) Execute(ctx context.Context, lead *entities.Lead, leadGroup *entities.LeadGroup) error {

	return uc.tx.Do(ctx, func(txCtx context.Context) error {

		leadCreateScope, err := uc.leadPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}

		if err := uc.leadRepo.Create(txCtx, lead, leadCreateScope); err != nil {
			return err
		}

		leadGroupScope, err := uc.leadGroupPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}

		leadGroup.Lead_id = lead.ID
		if err := uc.leadGroupsRepo.Create(txCtx, leadGroup, leadGroupScope); err != nil {
			return err
		}
		return nil
	})

}
