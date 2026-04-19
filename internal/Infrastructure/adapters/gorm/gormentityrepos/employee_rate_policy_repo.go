package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormEmployeeRatePoliciesRepo struct {
	*genericAdapter.GormRepository[entities.EmployeeRatePolicy]
	db *gorm.DB
}

func NewGormEmployeeRatePoliciesRepo(db *gorm.DB) entitiesrepos.EmployeeRatePolicyRepo {
	return &GormEmployeeRatePoliciesRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.EmployeeRatePolicy](db),
		db:             db,
	}
}
