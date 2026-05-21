package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormDanceStylesRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.DanceStyle]
}

func NewGormDanceStylesRepo(db *gorm.DB) entitiesrepos.DanceStylesRepo {
	return &GormDanceStylesRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.DanceStyle](db),
	}
}
