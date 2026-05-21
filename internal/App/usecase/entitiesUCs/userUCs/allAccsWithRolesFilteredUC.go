package userUCs

import (
	"context"

	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type GetAllAccountsWithRolesFilteredUC struct {
	userQueryService entitiesrepos.UserAccountQueryService
	policy           useraccountpolicies.IUserAccountReadPolicy
}

type IGetAllAccountsWithRolesFilteredUC interface {
	Execute(ctx context.Context, roles []string) ([]*businessobjects.UserBO, error)
}

func NewGetAllAccountsWithRolesFilteredUC(
	userQueryService entitiesrepos.UserAccountQueryService,
	policy useraccountpolicies.IUserAccountReadPolicy,
) IGetAllAccountsWithRolesFilteredUC {
	return &GetAllAccountsWithRolesFilteredUC{userQueryService, policy}
}

func (uc *GetAllAccountsWithRolesFilteredUC) Execute(ctx context.Context, roles []string) ([]*businessobjects.UserBO, error) {

	scope, err := uc.policy.CanReadAllEmpByRole(ctx, roles)
	if err != nil {
		return nil, err
	}

	userBO, err := uc.userQueryService.GetAllAccountsWithRoles(ctx, scope, roles...)
	if err != nil {
		return nil, err
	}

	return userBO, nil
}
