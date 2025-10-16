package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type SchoolRepository interface {
	generic.Repository[entities.School]
	Register(ctx context.Context, entity *entities.School) error
}
