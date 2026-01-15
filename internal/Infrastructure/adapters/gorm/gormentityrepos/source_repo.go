package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSourceRepository struct {
	*genericAdapter.GormRepository[entities.Source]
	db *gorm.DB
}

func NewGormSoutceRepository(db *gorm.DB) entitiesrepos.SourceRepo {
	return &GormSourceRepository{
		GormRepository: genericAdapter.NewGormRepository[entities.Source](db),
		db:             db,
	}
}
