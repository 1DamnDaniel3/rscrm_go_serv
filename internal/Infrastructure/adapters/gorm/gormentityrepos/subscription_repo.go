package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSubscriptionRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.Subscription]
}

func NewGormSubscriptionRepo(db *gorm.DB) entitiesrepos.SubscriptionRepo {
	return &GormSubscriptionRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.Subscription](db),
	}
}
