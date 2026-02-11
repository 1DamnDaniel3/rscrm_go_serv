package studentucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateStudentUC struct {
	tx            services.Transaction
	studRepo      entitiesrepos.StudentsRepo
	studGroupRepo entitiesrepos.StudentGroupsRepo
}

type ICreateStudentUC interface {
	Execute(ctx context.Context, student *entities.Student, studGroup *entities.StudentGroup) error
}

func NewCreateStudentUC(
	tx services.Transaction,
	studRepo entitiesrepos.StudentsRepo,
	studGroupRepo entitiesrepos.StudentGroupsRepo) *CreateStudentUC {
	return &CreateStudentUC{
		tx:            tx,
		studRepo:      studRepo,
		studGroupRepo: studGroupRepo,
	}
}

func (uc *CreateStudentUC) Execute(ctx context.Context, student *entities.Student, studGroup *entities.StudentGroup) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {
		if err := uc.studRepo.Create(txCtx, student); err != nil {
			return err
		}

		studGroup.Student_id = student.Id
		if err := uc.studGroupRepo.Create(txCtx, studGroup); err != nil {
			return err
		}
		return nil
	})
}
