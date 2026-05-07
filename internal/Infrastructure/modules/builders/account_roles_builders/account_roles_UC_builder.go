package accountrolesbuilders

import (
	accountrolespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/account_roles_policies"
	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	accountrolesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/account_roles_ucs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type AccountRolesUseCases struct {
	CRUD           genericcruduc.ICRUDUseCase[entities.AccountRoles]
	NotGenericCrud accountrolesucs.IAccountRolesCrudUC
}

func NewAccountRolesUseCases(
	accountRolesRepo entitiesrepos.AccountRolesRepo,
	accountRolesPolicies *accountrolespolicies.AccountRolesPolicies,

	userAccountRepo entitiesrepos.UserAccountRepository,
	userAccountPolicy *useraccountpolicies.UserAccountPolicies,
) *AccountRolesUseCases {

	return &AccountRolesUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			accountRolesRepo,
			accountRolesPolicies.CRUD,
		),
		NotGenericCrud: accountrolesucs.NewAccountRolesCrudUC(
			userAccountRepo,
			// userAccountPolicy.CRUD,

			accountRolesRepo,
			accountRolesPolicies.CreatePolicy,
			accountRolesPolicies.DeletePolicy,
		),
	}
}
