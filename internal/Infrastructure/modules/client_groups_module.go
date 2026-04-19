package modules

import (
	clientgroupspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_groups_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type ClientGroupsModule struct {
	ClientGroupRepo      entitiesrepos.ClientGroupRepo
	ClientGroupsPolicies *clientgroupspolicies.ClientGroupsPolicies
}

func NewClientGroupsModule(
	db *gorm.DB,
) *ClientGroupsModule {
	return &ClientGroupsModule{
		ClientGroupRepo: gormentityrepos.NewGormClientGroupRepo(db),

		ClientGroupsPolicies: clientgroupspolicies.NewClientGroupsPolicies(
			clientgroupspolicies.NewClientGroupsCrudPolicy(),
		),
	}
}
