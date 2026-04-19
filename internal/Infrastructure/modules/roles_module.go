package modules

import (
	rolepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/role_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type RolesModule struct {
	RolesRepo     entitiesrepos.RolesRepo
	RolesPolicies *rolepolicies.RolesPolicies
}

func NewRolesModule(
	db *gorm.DB,
) *RolesModule {
	return &RolesModule{
		RolesRepo: gormentityrepos.NewGormRolesRepo(db),

		RolesPolicies: rolepolicies.NewRolesPolicies(
			rolepolicies.NewRolesCrudPolicy(),
		),
	}
}
