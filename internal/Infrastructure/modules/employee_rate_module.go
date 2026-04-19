package modules

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type EmployeeRateModule struct {
	EmployeeRateRepo entitiesrepos.EmployeeRateRepo
}

func NewEmployeeRateModule(
	db *gorm.DB,
) *EmployeeRateModule {
	return &EmployeeRateModule{
		EmployeeRateRepo: gormentityrepos.NewGormEmployeeRateRepo(db),
	}
}
