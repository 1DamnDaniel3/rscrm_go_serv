package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormEmployeeRateRepo struct {
	*genericAdapter.GormRepository[entities.EmployeeRate]
	db *gorm.DB
}

func NewGormEmployeeRateRepo(db *gorm.DB) entitiesrepos.EmployeeRateRepo {
	return &GormEmployeeRateRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.EmployeeRate](db),
		db:             db,
	}
}
