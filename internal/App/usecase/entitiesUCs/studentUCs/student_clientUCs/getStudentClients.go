package studentclientUCs

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type GetStudentClientsUC struct {
	queryService entitiesrepos.StudentQueryService
}

type IGetStudentClientsUC interface {
	Execute(ctx context.Context, student_id int64, clientsSlice *[]businessobjects.GetStudentClientsBO) error
}

func NewGetStudentClientsUC(queryService entitiesrepos.StudentQueryService) *GetStudentClientsUC {
	return &GetStudentClientsUC{queryService}
}

func (uc *GetStudentClientsUC) Execute(ctx context.Context, student_id int64, clientsSlice *[]businessobjects.GetStudentClientsBO) error {
	return uc.queryService.GetStudentClients(ctx, student_id, clientsSlice)
}
