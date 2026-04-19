package studentgroupUCs

import (
	"context"

	studentgrouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_group_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type StudentGroupCRUDucs struct {
	repo   entitiesrepos.StudentGroupsRepo
	policy studentgrouppolicies.IStudentGroupCrudPolicy
}

type IStudentGroupCRUDucs interface {
	// Create(ctx context.Context, studentGroup *entities.StudentGroup) error
	Delete(ctx context.Context, stud_id int64, group_id int64) (*entities.StudentGroup, error)
}

func NewStudentGroupCRUDucs(
	repo entitiesrepos.StudentGroupsRepo,
	policy studentgrouppolicies.IStudentGroupCrudPolicy,
) IStudentGroupCRUDucs {
	return &StudentGroupCRUDucs{repo, policy}
}

// ---=== methods ===---

// func (uc *StudentGroupCRUDucs) Create(ctx context.Context, studentGroup *entities.StudentGroup) error {

// 	if err := uc.repo.Create(ctx, studentGroup); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (uc *StudentGroupCRUDucs) Delete(ctx context.Context, stud_id int64, group_id int64) (*entities.StudentGroup, error) {

	relationMap := map[string]any{
		"student_id": stud_id,
		"group_id":   group_id,
	}

	readScope, err := uc.policy.CanReadOne(ctx)
	if err != nil {
		return nil, err
	}
	relation, err := uc.repo.FindRelation(ctx, relationMap, readScope)
	if err != nil {
		return nil, err
	}

	delScope, err := uc.policy.CanDelete(ctx)
	if err != nil {
		return nil, err
	}
	if err := uc.repo.Delete(ctx, relation.ID, relation, delScope); err != nil {
		return nil, err
	}

	return relation, nil
}
