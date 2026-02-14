package clientucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateGroupedClientUC struct {
	tx             services.Transaction
	cli_groupsRepo entitiesrepos.ClientGroupRepo
	clientRepo     entitiesrepos.ClientRepo
}

type ICreateGroupedClientUC interface {
	Execute(ctx context.Context, lead *entities.Client, leadGroup *entities.ClientGroup) error
}

func NewCreateGroupedClientUC(
	tx services.Transaction,
	cli_groupsRepo entitiesrepos.ClientGroupRepo,
	clientRepo entitiesrepos.ClientRepo) *CreateGroupedClientUC {
	return &CreateGroupedClientUC{tx, cli_groupsRepo, clientRepo}
}

func (uc *CreateGroupedClientUC) Execute(ctx context.Context, client *entities.Client, clientGroup *entities.ClientGroup) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {
		if err := uc.clientRepo.Create(txCtx, client); err != nil {
			return err
		}
		clientGroup.Client_id = client.ID
		if err := uc.cli_groupsRepo.Create(txCtx, clientGroup); err != nil {
			return err
		}
		return nil
	})
}
