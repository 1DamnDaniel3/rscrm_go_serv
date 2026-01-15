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

type GormAccountRolesRepo struct {
	*genericAdapter.GormRepository[entities.AccountRoles]
	db *gorm.DB
}

func NewGormAccountRolesRepo(db *gorm.DB) entitiesrepos.AccountRolesRepo {
	return &GormAccountRolesRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.AccountRoles](db),
		db:             db,
	}
}

func (r *GormAccountRolesRepo) Register(ctx context.Context, entity *entities.AccountRoles) error {
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}
