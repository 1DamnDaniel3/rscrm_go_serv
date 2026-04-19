package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormEmployeeRateRuleRepo struct {
	*genericAdapter.GormRepository[entities.EmployeeRateRule]
	db *gorm.DB
}

func NewGormEmployeeRateRuleRepo(db *gorm.DB) entitiesrepos.EmployeeRateRuleRepo {
	return &GormEmployeeRateRuleRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.EmployeeRateRule](db),
		db:             db,
	}
}
