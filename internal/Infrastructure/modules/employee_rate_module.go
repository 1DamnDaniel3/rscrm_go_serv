package modules

import (
	employeeratepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/employee_rate_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type EmployeeRateModule struct {
	EmployeeRateRepo     entitiesrepos.EmployeeRateRepo
	EmployeeRatePolicies *employeeratepolicies.EmployeeRatePolicies
}

func NewEmployeeRateModule(
	db *gorm.DB,
) *EmployeeRateModule {
	return &EmployeeRateModule{
		EmployeeRateRepo: gormentityrepos.NewGormEmployeeRateRepo(db),
		EmployeeRatePolicies: employeeratepolicies.NewEmployeeRatePolicyPolicies(
			employeeratepolicies.NewEmployeeRatePolicyPolicy(),
		),
	}
}
