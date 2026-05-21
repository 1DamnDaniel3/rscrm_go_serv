package studentgroupUCs

import (
	"context"

	studentgrouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_group_policies"
	studentpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateStudentUC struct {
	tx            services.Transaction
	studRepo      entitiesrepos.StudentsRepo
	studGroupRepo entitiesrepos.StudentGroupsRepo

	studentPolicy      studentpolicies.IStudentCrudPolicy
	studentGroupPolicy studentgrouppolicies.IStudentGroupCrudPolicy
}

type ICreateStudentUC interface {
	Execute(ctx context.Context, student *entities.Student, studGroup *entities.StudentGroup) error
}

func NewCreateStudentUC(
	tx services.Transaction,
	studRepo entitiesrepos.StudentsRepo,
	studGroupRepo entitiesrepos.StudentGroupsRepo,

	studentPolicy studentpolicies.IStudentCrudPolicy,
	studentGroupPolicy studentgrouppolicies.IStudentGroupCrudPolicy,
) *CreateStudentUC {
	return &CreateStudentUC{
		tx:            tx,
		studRepo:      studRepo,
		studGroupRepo: studGroupRepo,

		studentPolicy:      studentPolicy,
		studentGroupPolicy: studentGroupPolicy,
	}
}

func (uc *CreateStudentUC) Execute(ctx context.Context, student *entities.Student, studGroup *entities.StudentGroup) error {
	return uc.tx.Do(ctx, func(txCtx context.Context) error {

		//stud policy
		studScope, err := uc.studentPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}
		//stud repo
		if err := uc.studRepo.Create(txCtx, student, studScope); err != nil {
			return err
		}

		// stud_group policy
		studGroupPolicy, err := uc.studentGroupPolicy.CanCreate(txCtx)
		if err != nil {
			return err
		}

		//stud_group repo
		studGroup.Student_id = student.ID
		if err := uc.studGroupRepo.Create(txCtx, studGroup, studGroupPolicy); err != nil {
			return err
		}
		return nil
	})
}
