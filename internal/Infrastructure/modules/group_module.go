package modules

import (
	grouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/group_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type GroupModule struct {
	GroupRepo     entitiesrepos.GroupRepo
	GroupPolicies *grouppolicies.GroupPolicies
}

func NewGroupModule(
	db *gorm.DB,
) *GroupModule {
	return &GroupModule{
		GroupRepo: gormentityrepos.NewGormGroupRepository(db),

		GroupPolicies: grouppolicies.NewGroupPolicies(
			grouppolicies.NewGroupCrudPolicy(),
		),
	}
}
