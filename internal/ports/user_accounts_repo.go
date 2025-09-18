package ports

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
)

type UserAccountRepository interface {
	GetByEmail(email string) (*entities.UserAccount, error)
}
