package leadUCs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateLeadUC struct {
	tx             services.Transaction
	leadRepo       entitiesrepos.LeadsRepository
	leadGroupsRepo entitiesrepos.LeadGroupsRepo
}

type ICreateLeadUC interface {
	Execute(ctx context.Context, lead *entities.Lead, leadGroup *entities.LeadGroup) error
}

func NewCreateLeadUC(
	tx services.Transaction,
	leadRepo entitiesrepos.LeadsRepository,
	leadGroupsRepo entitiesrepos.LeadGroupsRepo) *CreateLeadUC {
	return &CreateLeadUC{
		tx:             tx,
		leadRepo:       leadRepo,
		leadGroupsRepo: leadGroupsRepo,
	}
}

func (uc *CreateLeadUC) Execute(ctx context.Context, lead *entities.Lead, leadGroup *entities.LeadGroup) error {

	return uc.tx.Do(ctx, func(txCtx context.Context) error {
		if err := uc.leadRepo.Create(txCtx, lead); err != nil {
			return err
		}
		leadGroup.Lead_id = lead.ID
		if err := uc.leadGroupsRepo.Create(txCtx, leadGroup); err != nil {
			return err
		}
		return nil
	})

}
