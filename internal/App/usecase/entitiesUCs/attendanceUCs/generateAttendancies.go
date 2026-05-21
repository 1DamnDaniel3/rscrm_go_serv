package attendanceucs

import (
	"context"
	"fmt"

	attendancepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/attendance_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type GenerateAttendanceUC struct {
	tx                     services.Transaction
	AttendanceRepo         entitiesrepos.AttendanceRepo
	AttendanceQueryService entitiesrepos.AttendanceQueryService
	CreatePolicy           attendancepolicies.IAttendanceCreatePolicy

	GetGroupedStudentsUC studentucs.IGroupedStudentsUC
}

type IGenerateAttendanceUC interface {
	Execute(ctx context.Context, group_id int64, lesson_id int64) ([]entities.Attendance, error)
}

func NewGenerateAttendanceUC(
	tx services.Transaction,
	AttendanceRepo entitiesrepos.AttendanceRepo,
	AttendanceQueryService entitiesrepos.AttendanceQueryService,

	CreatePolicy attendancepolicies.IAttendanceCreatePolicy,
	GetGroupedStudentsUC studentucs.IGroupedStudentsUC,
) IGenerateAttendanceUC {
	return &GenerateAttendanceUC{
		tx,
		AttendanceRepo,
		AttendanceQueryService,
		CreatePolicy,
		GetGroupedStudentsUC,
	}
}

func (uc *GenerateAttendanceUC) Execute(
	ctx context.Context,
	group_id int64,
	lesson_id int64,
) ([]entities.Attendance, error) {

	var attendances []entities.Attendance

	err := uc.tx.Do(ctx, func(txCtx context.Context) error {

		scope, err := uc.CreatePolicy.CanGenerateAttendance(ctx)
		if err != nil {
			return err
		}

		students := &[]entities.Student{}
		if err := uc.GetGroupedStudentsUC.Execute(txCtx, group_id, students); err != nil {
			return err
		}

		if len(*students) == 0 {
			attendances = []entities.Attendance{}
			return fmt.Errorf("empty group") // handler проверяет этот err.Error()
		}

		student_ids := make([]int64, len(*students))
		attendanciesToCreate := make([]entities.Attendance, len(*students))
		for i, student := range *students {
			student_ids[i] = student.ID
			attendanciesToCreate[i] = entities.Attendance{
				Student_id: student.ID,
				Lesson_id:  lesson_id,
				Status:     "absent",
				Marked_at:  nil,
			}
		}

		if err := uc.AttendanceRepo.CreateMany(txCtx, &attendanciesToCreate, scope); err != nil {
			return err
		}

		attendances, err = uc.AttendanceQueryService.FindAttendancies(txCtx, student_ids, lesson_id, nil)
		if err != nil {
			return err
		}

		return nil

	})
	if err != nil {
		return nil, err
	}

	return attendances, nil
}
