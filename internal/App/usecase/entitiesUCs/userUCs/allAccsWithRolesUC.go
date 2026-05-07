package userUCs

import (
	"context"

	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type GetAllAccountsWithRolesUC struct {
	userQueryService entitiesrepos.UserAccountQueryService
	policy           useraccountpolicies.IUserAccountCrudPolicy
}

type IGetAllAccountsWithRolesUC interface {
	Execute(ctx context.Context) ([]*businessobjects.UserBO, error)
}

func NewGetAllAccountsWithRolesUC(
	userQueryService entitiesrepos.UserAccountQueryService,
	policy useraccountpolicies.IUserAccountCrudPolicy,
) IGetAllAccountsWithRolesUC {
	return &GetAllAccountsWithRolesUC{userQueryService, policy}
}

func (uc *GetAllAccountsWithRolesUC) Execute(ctx context.Context) ([]*businessobjects.UserBO, error) {

	scope, err := uc.policy.CanReadAll(ctx)
	if err != nil {
		return nil, err
	}

	userBO, err := uc.userQueryService.GetAllAccountsWithRoles(ctx, scope)
	if err != nil {
		return nil, err
	}

	return userBO, nil
}
