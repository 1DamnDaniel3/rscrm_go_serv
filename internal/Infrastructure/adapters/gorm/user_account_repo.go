package adapters

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormUserAccountRepo struct {
	*genericAdapter.GormRepository[entities.UserAccount]
	db     *gorm.DB
	hasher ports.PasswordHasher
}

func (r *GormUserAccountRepo) Create(entity *entities.UserAccount) error {
	var err error
	entity.Password, err = r.hasher.Hash(entity.Password)
	if err != nil {
		return err
	}
	return r.db.Create(entity).Error
}

func (r *GormUserAccountRepo) GetByEmail(email string) (*entities.UserAccount, error) {
	var user entities.UserAccount
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewGormUserAccountRepo(db *gorm.DB, hasher ports.PasswordHasher) ports.UserAccountRepository {
	return &GormUserAccountRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.UserAccount](db),
		db:             db,
		hasher:         hasher}
}
