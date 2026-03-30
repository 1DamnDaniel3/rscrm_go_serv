package studentucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetStudentGroupUC struct {
	studentQueryService entitiesrepos.StudentQueryService
}

type IGetStudentGroupUC interface {
	Execute(ctx context.Context, student_id int64, groupSlice *[]entities.Group) error
}

func NewGetStudentGroupUC(studentQueryService entitiesrepos.StudentQueryService) IGetStudentGroupUC {
	return &GetStudentGroupUC{studentQueryService}
}

func (uc *GetStudentGroupUC) Execute(ctx context.Context, student_id int64, groupSlice *[]entities.Group) error {
	return uc.studentQueryService.GetStudentGroups(ctx, student_id, groupSlice)
}
