package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type UserAccountRepository interface {
	genericrepo.Repository[entities.UserAccount]
	Create(ctx context.Context, entity *entities.UserAccount, scope *policytypes.Scope) error
	GetByEmail(email string) (*entities.UserAccount, error)
	Register(ctx context.Context, entity *entities.UserAccount) error
}

type UserAccountQueryService interface {
	GetMe(ctx context.Context) (*businessobjects.UserBO, error)
	GetAllAccountsWithRoles(ctx context.Context, scope *policytypes.Scope) ([]*businessobjects.UserBO, error)
}
