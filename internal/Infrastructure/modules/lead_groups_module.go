package modules

import (
	leadgrouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lead_group_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type LeadGroupsModule struct {
	LeadGroupsRepo    entitiesrepos.LeadGroupsRepo
	LeadGroupPolicies *leadgrouppolicies.LeadGroupPolicies
}

func NewLeadGroupsModule(
	db *gorm.DB,
) *LeadGroupsModule {
	return &LeadGroupsModule{
		LeadGroupsRepo: gormentityrepos.NewGormLeadGroupsRepo(db),

		LeadGroupPolicies: leadgrouppolicies.NewLeadGroupPolicies(
			leadgrouppolicies.NewLeadGroupCrudPolicy(),
		),
	}
}
