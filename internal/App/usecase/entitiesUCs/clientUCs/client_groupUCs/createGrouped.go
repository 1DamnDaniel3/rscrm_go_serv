package clientgroupsUCs

import (
	"context"

	clientgroupspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_groups_policies"
	clientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateGroupedClientUC struct {
	tx             services.Transaction
	cli_groupsRepo entitiesrepos.ClientGroupRepo
	clientRepo     entitiesrepos.ClientRepo

	clientGroupsPolicy clientgroupspolicies.IClientGroupsCrudPolicy
	clientPolicy       clientpolicies.IClientCrudPolicy
}

type ICreateGroupedClientUC interface {
	Execute(ctx context.Context, lead *entities.Client, leadGroup *entities.ClientGroup) error
}

func NewCreateGroupedClientUC(
	tx services.Transaction,
	cli_groupsRepo entitiesrepos.ClientGroupRepo,
	clientRepo entitiesrepos.ClientRepo,

	clientGroupsPolicy clientgroupspolicies.IClientGroupsCrudPolicy,
	clientPolicy clientpolicies.IClientCrudPolicy,
) *CreateGroupedClientUC {
	return &CreateGroupedClientUC{tx, cli_groupsRepo, clientRepo, clientGroupsPolicy, clientPolicy}
}

func (uc *CreateGroupedClientUC) Execute(ctx context.Context, client *entities.Client, clientGroup *entities.ClientGroup) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {

		clientCreateScope, err := uc.clientPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}

		if err := uc.clientRepo.Create(txCtx, client, clientCreateScope); err != nil {
			return err
		}

		clientGroupCreateScope, err := uc.clientGroupsPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}

		clientGroup.Client_id = client.ID
		if err := uc.cli_groupsRepo.Create(txCtx, clientGroup, clientGroupCreateScope); err != nil {
			return err
		}
		return nil
	})
}
