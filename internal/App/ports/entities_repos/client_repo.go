package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type ClientRepo interface {
	generic.Repository[entities.Client]
	GetGroupedClients(ctx context.Context, group_id int64, entities *[]entities.Client) error
}

type ClientsQueryService interface {
	// GetClientStudents(ctx context.Context, student_id int64, clientsSlice *[]businessobjects.GetStudentClientsBO) error
	Search(ctx context.Context, input string, clientSlice *[]entities.Client) error
}
