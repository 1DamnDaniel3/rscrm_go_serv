package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type SchoolRepository interface {
	genericrepo.Repository[entities.School]
	Register(ctx context.Context, entity *entities.School) error
}
