package clientstudentsUCs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetClientStudentsUC struct {
	queryService entitiesrepos.ClientsQueryService
}

type IGetClientStudentsUC interface {
	Execute(ctx context.Context, client_id int64, studentSlice *[]entities.Student) error
}

func NewGetClientStudentsUC(queryService entitiesrepos.ClientsQueryService) *GetClientStudentsUC {
	return &GetClientStudentsUC{queryService}
}

func (uc *GetClientStudentsUC) Execute(ctx context.Context,
	client_id int64, studentSlice *[]entities.Student) error {
	return uc.queryService.GetClientStudents(ctx, client_id, studentSlice)
}
