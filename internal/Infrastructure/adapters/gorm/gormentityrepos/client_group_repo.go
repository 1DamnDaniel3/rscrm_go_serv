package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormClientGroupRepo struct {
	*genericAdapter.GormRepository[entities.ClientGroup]
	db *gorm.DB
}

func NewGormClientGroupRepo(db *gorm.DB) entitiesrepos.ClientGroupRepo {
	return &GormClientGroupRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.ClientGroup](db),
		db:             db,
	}
}
