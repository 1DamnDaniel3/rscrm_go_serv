package gormentityrepos

import (
	"context"
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormUserAccountRepo struct {
	*genericAdapter.GormRepository[entities.UserAccount]
	db     *gorm.DB
	hasher ports.PasswordHasher
}

func NewGormUserAccountRepo(db *gorm.DB, hasher ports.PasswordHasher) entitiesrepos.UserAccountRepository {
	return &GormUserAccountRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.UserAccount](db),
		db:             db,
		hasher:         hasher}
}

func (r *GormUserAccountRepo) Register(ctx context.Context, entity *entities.UserAccount) error {
	var err error
	entity.Password, err = r.hasher.Hash(entity.Password)
	if err != nil {
		return err
	}
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func (r *GormUserAccountRepo) GetByEmail(email string) (*entities.UserAccount, error) {
	var user entities.UserAccount
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ======================================== OVERRIDE =============================================

func (r *GormUserAccountRepo) Create(
	ctx context.Context,
	entity *entities.UserAccount,
	scope *policytypes.Scope,
) error {

	db := gormutils.DBFromCtx(ctx, r.db)

	// scope
	if scope != nil && !scope.IsGlobal {
		entity.School_id = scope.School_id
	}

	var err error
	entity.Password, err = r.hasher.Hash(entity.Password)
	if err != nil {
		return err
	}

	return db.Create(entity).Error
}
