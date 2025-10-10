package ports

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports/generic"
)

type UserAccountRepository interface {
	generic.Repository[entities.UserAccount]
	Create(entity *entities.UserAccount) error
	GetByEmail(email string) (*entities.UserAccount, error)
}
