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
	Delete(ctx context.Context, id int64, studentGroup *entities.StudentGroup) error
}

func NewCreateStudentGroupRelationUC(repo entitiesrepos.StudentGroupsRepo) IStudentGroupCRUDucs {
	return &StudentGroupCRUDucs{repo}
}

// ---=== methods ===---

func (uc *StudentGroupCRUDucs) Create(ctx context.Context, studentGroup *entities.StudentGroup) error {

	if err := uc.repo.Create(ctx, studentGroup); err != nil {
		return err
	}

	return nil
}

func (uc *StudentGroupCRUDucs) Delete(ctx context.Context, id int64, studentGroup *entities.StudentGroup) error {
	if err := uc.repo.Delete(ctx, id, studentGroup); err != nil {
		return err
	}

	return nil
}
