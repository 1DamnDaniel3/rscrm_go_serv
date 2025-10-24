package adapters

import (
	"context"
	"errors"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormSchoolRepo struct {
	*genericAdapter.GormRepository[entities.School]
	db *gorm.DB
}

func (r *GormSchoolRepo) Register(ctx context.Context, entity *entities.School) error {
	tx, ok := ctx.Value(txKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func NewGormSchoolRepo(db *gorm.DB) entitiesrepos.SchoolRepository {
	return &GormSchoolRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.School](db),
		db:             db,
	}
}
