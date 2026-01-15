package gormentityrepos

import (
	"context"
	"errors"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormUserProfileRepo struct {
	*genericAdapter.GormRepository[entities.UserProfile]
	db *gorm.DB
}

func (r *GormUserProfileRepo) Register(ctx context.Context, entity *entities.UserProfile) error {
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func NewGormUserProfileRepo(db *gorm.DB) entitiesrepos.ProfileRepo {
	if db == nil {
		panic("NewGormUserProfileRepo: db is nil — database not initialized")
	}
	return &GormUserProfileRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.UserProfile](db),
		db:             db,
	}
}
