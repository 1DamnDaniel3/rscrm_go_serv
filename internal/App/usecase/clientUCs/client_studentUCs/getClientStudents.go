package clientstudentsUCs

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type GetClientStudentsUC struct {
	queryService entitiesrepos.ClientsQueryService
}

type IGetClientStudentsUC interface {
	Execute(ctx context.Context, client_id int64, studentSlice *[]businessobjects.GetClientStudentsBO) error
}

func NewGetClientStudentsUC(queryService entitiesrepos.ClientsQueryService) IGetClientStudentsUC {
	return &GetClientStudentsUC{queryService}
}

func (uc *GetClientStudentsUC) Execute(ctx context.Context,
	client_id int64, studentSlice *[]businessobjects.GetClientStudentsBO) error {
	return uc.queryService.GetClientStudents(ctx, client_id, studentSlice)
}
