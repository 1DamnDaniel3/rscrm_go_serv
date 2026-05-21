package modules

import (
	salaryaccuralitemspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/salary_accural_items_policies.go"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type SalaryAccuralItemsModule struct {
	SalaryAccuralItemsRepo     entitiesrepos.SalaryAccuralItemsRepo
	SalaryAccuralItemsPolicies *salaryaccuralitemspolicies.SalaryAccuralItemsPolicies
}

func NewSalaryAccuralItemsModule(db *gorm.DB) *SalaryAccuralItemsModule {
	return &SalaryAccuralItemsModule{
		SalaryAccuralItemsRepo: gormentityrepos.NewGormSalaryAccuralItemsRepo(db),
		SalaryAccuralItemsPolicies: salaryaccuralitemspolicies.NewSalaryAccuralItemsPolicies(
			salaryaccuralitemspolicies.NewSalaryAccuralItemsCrudPolicy(),
		),
	}
}
