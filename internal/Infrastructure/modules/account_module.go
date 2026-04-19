package modules

import (
	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type AccountModule struct {
	UserRepo        entitiesrepos.UserAccountRepository
	AccountPolicies *useraccountpolicies.UserAccountPolicies
}

func NewAccountModule(
	db *gorm.DB,
	hasher ports.PasswordHasher,
) *AccountModule {
	return &AccountModule{
		// -=== repo
		UserRepo: gormentityrepos.NewGormUserAccountRepo(db, hasher),

		// -=== policies
		AccountPolicies: useraccountpolicies.NewUserAccountPolicies(
			useraccountpolicies.NewUserAccountCrudPolicy(),
		),
	}
}
