package adapters

import (
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports/generic"
	"gorm.io/gorm"
)

type GormUserAccountRepo struct {
	generic.Repository[entities.UserAccount]
	db *gorm.DB
}

func (r *GormUserAccountRepo) GetByEmail(email string) (*entities.UserAccount, error) {
	var user entities.UserAccount
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func NewGormUserAccountRepo(db *gorm.DB) ports.UserAccountRepository {
	return &GormUserAccountRepo{
		Repository: genericAdapter.NewGormRepository[entities.UserAccount](db),
		db:         db}
}
