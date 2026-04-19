package modules

import (
	accountrolespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/account_roles_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type AccountRolesModule struct {
	AccountRolesRepo    entitiesrepos.AccountRolesRepo
	AccountRolePolicies *accountrolespolicies.AccountRolesPolicies
}

func NewAccountRolesModule(
	db *gorm.DB,
) *AccountRolesModule {
	return &AccountRolesModule{
		AccountRolesRepo: gormentityrepos.NewGormAccountRolesRepo(db),

		AccountRolePolicies: accountrolespolicies.NewAccountRolesPolicies(
			accountrolespolicies.NewAccountRolesCrudPolicy(),
		),
	}
}
