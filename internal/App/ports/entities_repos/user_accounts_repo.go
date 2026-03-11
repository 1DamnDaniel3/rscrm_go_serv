package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type UserAccountRepository interface {
	genericrepo.Repository[entities.UserAccount]
	Create(ctx context.Context, entity *entities.UserAccount) error
	GetByEmail(email string) (*entities.UserAccount, error)
	Register(ctx context.Context, entity *entities.UserAccount) error
}
