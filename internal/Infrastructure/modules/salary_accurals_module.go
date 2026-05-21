package modules

import (
	salaryaccuralspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/salary_accruals_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type SalaryAccuralsModule struct {
	SalaryAccuralsRepo     entitiesrepos.SalaryAccuralsRepo
	SalaryAccuralsPolicies *salaryaccuralspolicies.SalaryAccuralsPolicies
}

func NewSalaryAccuralsModule(db *gorm.DB) *SalaryAccuralsModule {
	return &SalaryAccuralsModule{
		SalaryAccuralsRepo: gormentityrepos.NewGormSalaryAccuralsRepo(db),
		SalaryAccuralsPolicies: salaryaccuralspolicies.NewSalaryAccuralsPolicies(
			salaryaccuralspolicies.NewSalaryAccuralsCrudPolicy(),
		),
	}
}
