package modules

import (
	leadpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lead_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type LeadModule struct {
	LeadRepo     entitiesrepos.LeadsRepository
	LeadPolicies *leadpolicies.LeadPolicies
}

func NewLeadModule(
	db *gorm.DB,
) *LeadModule {
	return &LeadModule{
		LeadRepo: gormentityrepos.NewGormLeadsRepo(db),

		LeadPolicies: leadpolicies.NewLeadPolicies(
			leadpolicies.NewLeadCrudPolicy(),
		),
	}
}
