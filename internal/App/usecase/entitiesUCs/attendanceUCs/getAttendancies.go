package attendanceucs

import (
	"context"

	attendancepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/attendance_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetAttendanciesUC struct {
	repo entitiesrepos.AttendanceRepo

	generateUC IGenerateAttendanceUC

	policy attendancepolicies.IAttendanceCrudPolicy
}

type IGetAttendanciesUC interface {
	Execute(ctx context.Context, group_id, lesson_id int64) ([]entities.Attendance, error)
}

func NewGetAttendanciesUC(
	repo entitiesrepos.AttendanceRepo,

	generateUC IGenerateAttendanceUC,

	policy attendancepolicies.IAttendanceCrudPolicy,
) IGetAttendanciesUC {
	return &GetAttendanciesUC{
		repo, generateUC, policy,
	}
}

// USE WHEN MADE UC FOR ARCHIVE/CLEAR OLD ATTENDANCIES
func (uc *GetAttendanciesUC) Execute(ctx context.Context, group_id, lesson_id int64) ([]entities.Attendance, error) {

	// readScope, err := uc.policy.CanReadAll(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// can use Redis cache to make conditions - generate attendancies below or not.

	// Generate attendancies
	generatedAttendancies, err := uc.generateUC.Execute(ctx, group_id, lesson_id)
	if err != nil {
		return nil, err
	}

	return generatedAttendancies, nil
}
