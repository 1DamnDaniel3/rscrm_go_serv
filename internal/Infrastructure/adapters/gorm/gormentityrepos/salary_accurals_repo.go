package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSalaryAccuralsRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.SalaryAccural]
}

func NewGormSalaryAccuralsRepo(db *gorm.DB) entitiesrepos.SalaryAccuralsRepo {
	return &GormSalaryAccuralsRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.SalaryAccural](db),
	}
}
