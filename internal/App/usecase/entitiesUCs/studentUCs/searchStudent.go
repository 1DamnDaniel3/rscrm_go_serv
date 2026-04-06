package studentucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type SearchStudentsUC struct {
	queryService entitiesrepos.StudentQueryService
}

type ISearchStudentsUC interface {
	Execute(ctx context.Context, input string, studentSlice *[]entities.Student) error
}

func NewSearchStudentsUC(queryService entitiesrepos.StudentQueryService) ISearchStudentsUC {
	return &SearchStudentsUC{queryService}
}

func (uc *SearchStudentsUC) Execute(ctx context.Context, input string, studentSlice *[]entities.Student) error {
	return uc.queryService.Search(ctx, input, studentSlice)
}
