package adapters

import (
	"context"
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSchoolRepo struct {
	*genericAdapter.GormRepository[entities.School]
}

func (r *GormSchoolRepo) Register(ctx context.Context, entity *entities.School) error {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func NewGormSchoolRepo() *GormSchoolRepo {
	return &GormSchoolRepo{}
}
