package ports

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type UserAccountRepository interface {
	generic.Repository[entities.UserAccount]
	Create(entity *entities.UserAccount) error
	GetByEmail(email string) (*entities.UserAccount, error)
}
