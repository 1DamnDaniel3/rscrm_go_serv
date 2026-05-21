package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSalaryAccuralItemsRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.SalaryAccrualItems]
}

func NewGormSalaryAccuralItemsRepo(db *gorm.DB) entitiesrepos.SalaryAccuralItemsRepo {
	return &GormSalaryAccuralItemsRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.SalaryAccrualItems](db),
	}
}
