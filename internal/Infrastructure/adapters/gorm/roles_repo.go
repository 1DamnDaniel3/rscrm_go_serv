package adapters

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormRolesRepo struct {
	*genericAdapter.GormRepository[entities.Roles]
	db *gorm.DB
}

func NewGormRolesRepo(db *gorm.DB) entitiesrepos.RolesRepo {
	return &GormRolesRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Roles](db),
		db:             db,
	}
}
