package studentgroupUCs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type StudentGroupCRUDucs struct {
	repo entitiesrepos.StudentGroupsRepo
}

type IStudentGroupCRUDucs interface {
	Create(ctx context.Context, studentGroup *entities.StudentGroup) error
	Delete(ctx context.Context, stud_id int64, group_id int64) (*entities.StudentGroup, error)
}

func NewStudentGroupCRUDucs(repo entitiesrepos.StudentGroupsRepo) IStudentGroupCRUDucs {
	return &StudentGroupCRUDucs{repo}
}

// ---=== methods ===---

func (uc *StudentGroupCRUDucs) Create(ctx context.Context, studentGroup *entities.StudentGroup) error {

	if err := uc.repo.Create(ctx, studentGroup); err != nil {
		return err
	}

	return nil
}

func (uc *StudentGroupCRUDucs) Delete(ctx context.Context, stud_id int64, group_id int64) (*entities.StudentGroup, error) {

	relationMap := map[string]any{
		"student_id": stud_id,
		"group_id":   group_id,
	}
	relation, err := uc.repo.FindRelation(ctx, relationMap)
	if err != nil {
		return nil, err
	}

	if err := uc.repo.Delete(ctx, relation.ID, relation); err != nil {
		return nil, err
	}

	return relation, nil
}
