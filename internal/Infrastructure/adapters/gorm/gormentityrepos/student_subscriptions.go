package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormStudentSubscriptionsRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.StudentSubscription]
}

func NewGormStudentSubscriptionsRepo(db *gorm.DB) entitiesrepos.StudentSubscriptionsRepo {
	return &GormStudentSubscriptionsRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.StudentSubscription](db),
	}
}
