package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormStatusRepository struct {
	*genericAdapter.GormRepository[entities.Status]
	db *gorm.DB
}

func NewGormStatusRepository(db *gorm.DB) entitiesrepos.StatusRepo {
	return &GormStatusRepository{
		GormRepository: genericAdapter.NewGormRepository[entities.Status](db),
		db:             db,
	}
}
