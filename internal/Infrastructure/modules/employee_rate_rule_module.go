package modules

import (
	employeeraterulespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/employee_rate_rules_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type EmployeeRateRuleModule struct {
	EmployeeRateRuleRepo     entitiesrepos.EmployeeRateRuleRepo
	EmployeeRateRulePolicies *employeeraterulespolicies.EmployeeRateRulesPolicies
}

func NewEmployeeRateRuleModule(
	db *gorm.DB,
) *EmployeeRateRuleModule {
	return &EmployeeRateRuleModule{
		EmployeeRateRuleRepo: gormentityrepos.NewGormEmployeeRateRuleRepo(db),
		EmployeeRateRulePolicies: employeeraterulespolicies.NewEmployeeRateRulesPolicies(
			employeeraterulespolicies.NewEmployeeRateRulesCrudPolicy(),
		),
	}
}
