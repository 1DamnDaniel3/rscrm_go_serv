package studentucs

import (
	"context"

	studentpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupedStudentsUC struct {
	repo   entitiesrepos.StudentsRepo
	policy studentpolicies.IStudentCrudPolicy
}

type IGroupedStudentsUC interface {
	Execute(ctx context.Context, group_id int64, entities *[]entities.Student) error
}

func NewGroupedStudentsUC(
	repo entitiesrepos.StudentsRepo,
	policy studentpolicies.IStudentCrudPolicy,
) *GroupedStudentsUC {
	return &GroupedStudentsUC{repo, policy}
}

func (uc *GroupedStudentsUC) Execute(ctx context.Context, group_id int64, entities *[]entities.Student) error {

	scope, err := uc.policy.CanReadAll(ctx)
	if err != nil {
		return err
	}

	return uc.repo.GetGroupedStudents(ctx, group_id, entities, scope)
}
