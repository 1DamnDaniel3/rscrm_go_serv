package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type AccountRolesRepo interface {
	genericrepo.Repository[entities.AccountRoles]
	Register(ctx context.Context, entity *entities.AccountRoles) error
}
