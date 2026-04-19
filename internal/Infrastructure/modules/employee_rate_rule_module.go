package modules

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type EmployeeRateRuleModule struct {
	EmployeeRateRuleRepo entitiesrepos.EmployeeRateRuleRepo
}

func NewEmployeeRateRuleModule(
	db *gorm.DB,
) *EmployeeRateRuleModule {
	return &EmployeeRateRuleModule{
		EmployeeRateRuleRepo: gormentityrepos.NewGormEmployeeRateRuleRepo(db),
	}
}
