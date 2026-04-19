package modules

import (
	employeeratepolicypolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/employee_rate_policy_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type EmployeeRatePolicyModule struct {
	EmployeeRatePolicyRepo     entitiesrepos.EmployeeRatePolicyRepo
	EmployeeRatePolicyPolicies *employeeratepolicypolicies.EmployeeRatePolicyPolicies
}

func NewEmployeeRatePolicyModule(
	db *gorm.DB,
) *EmployeeRatePolicyModule {
	return &EmployeeRatePolicyModule{
		EmployeeRatePolicyRepo: gormentityrepos.NewGormEmployeeRatePoliciesRepo(db),

		EmployeeRatePolicyPolicies: employeeratepolicypolicies.NewEmployeeRatePolicyPolicies(
			employeeratepolicypolicies.NewEmployeeRatePolicyPolicy(),
		),
	}
}
