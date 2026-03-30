package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormScheduleRepo struct {
	*genericAdapter.GormRepository[entities.Schedule]
	db *gorm.DB
}

func NewGormScheduleRepo(db *gorm.DB) entitiesrepos.ScheduleRepo {
	return &GormScheduleRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Schedule](db),
		db:             db,
	}
}
