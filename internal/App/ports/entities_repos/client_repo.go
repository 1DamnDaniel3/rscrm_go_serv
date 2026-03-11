package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type ClientRepo interface {
	genericrepo.Repository[entities.Client]
	GetGroupedClients(ctx context.Context, group_id int64, entities *[]entities.Client) error
}

type ClientsQueryService interface {
	GetClientStudents(ctx context.Context, client_id int64, studentsSlice *[]entities.Student) error
	GetClientGroups(ctx context.Context, client_id int64, groupSlice *[]entities.Group) error
	Search(ctx context.Context, input string, clientSlice *[]entities.Client) error
}
