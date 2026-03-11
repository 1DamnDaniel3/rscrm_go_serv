package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type LeadsRepository interface {
	genericrepo.Repository[entities.Lead]
	GetGroupedLeads(ctx context.Context, group_id int64, entities *[]entities.Lead) error
}
