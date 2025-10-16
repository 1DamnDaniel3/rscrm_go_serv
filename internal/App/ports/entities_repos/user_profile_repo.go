package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type ProfileRepo interface {
	generic.Repository[entities.UserProfile]
	Register(ctx context.Context, entity *entities.UserProfile) error
}
