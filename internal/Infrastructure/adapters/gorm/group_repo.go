package adapters

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormGroupRepository struct {
	*genericAdapter.GormRepository[entities.Group]
	db *gorm.DB
}

func NewGormGroupRepository(db *gorm.DB) entitiesrepos.GroupRepo {
	return &GormGroupRepository{
		GormRepository: genericAdapter.NewGormRepository[entities.Group](db),
		db:             db,
	}
}
