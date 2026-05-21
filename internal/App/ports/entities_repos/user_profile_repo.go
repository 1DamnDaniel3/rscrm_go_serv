package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type UserProfileRepo interface {
	genericrepo.Repository[entities.UserProfile]
	Register(ctx context.Context, entity *entities.UserProfile) error

	GetAllProfilesByRoles(ctx context.Context, scope *policytypes.Scope, roles ...string) (*[]entities.UserProfile, error)
}
