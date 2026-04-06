package studentucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupedStudentsUC struct {
	repo entitiesrepos.StudentsRepo
}

type IGroupedStudentsUC interface {
	Execute(ctx context.Context, group_id int64, entities *[]entities.Student) error
}

func NewGroupedStudentsUC(repo entitiesrepos.StudentsRepo) *GroupedStudentsUC {
	return &GroupedStudentsUC{repo}
}

func (r *GroupedStudentsUC) Execute(ctx context.Context, group_id int64, entities *[]entities.Student) error {
	return r.repo.GetGroupedStudents(ctx, group_id, entities)
}
