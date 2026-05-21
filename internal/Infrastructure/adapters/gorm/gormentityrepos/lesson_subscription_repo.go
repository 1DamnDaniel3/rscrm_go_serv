package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormLessonSubscriptionRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.LessonSubscription]
}

func NewGormLessonSubscriptionRepo(
	db *gorm.DB,
) entitiesrepos.LessonSubscriptionRepo {
	return &GormLessonSubscriptionRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.LessonSubscription](db),
	}
}
