package adapters

import (
	"context"
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormUserProfileRepo struct {
	*generic.GormRepository[entities.UserProfile]
	// db *gorm.DB
}

func (r *GormUserProfileRepo) Register(ctx context.Context, entity *entities.UserProfile) error {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func NewGormUserProfileRepo() *GormUserProfileRepo {
	return &GormUserProfileRepo{}
}
