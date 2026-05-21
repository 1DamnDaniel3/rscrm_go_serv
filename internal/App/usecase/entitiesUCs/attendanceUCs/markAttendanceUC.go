package attendanceucs

import (
	"context"

	attendancepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/attendance_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type MarkAttendanceUC struct {
	repo         entitiesrepos.AttendanceRepo
	updatePolicy attendancepolicies.IAttendanceUpdatePolicy
}

type IMarkAttendanceUC interface {
	Execute(ctx context.Context, attenance_id int64) (*entities.Attendance, error)
}

func NewMarkAttendanceUC(
	repo entitiesrepos.AttendanceRepo,
	updatePolicy attendancepolicies.IAttendanceUpdatePolicy,
) IMarkAttendanceUC {
	return &MarkAttendanceUC{
		repo, updatePolicy,
	}
}

func (uc *MarkAttendanceUC) Execute(ctx context.Context, attenance_id int64) (*entities.Attendance, error) {

	scope, err := uc.updatePolicy.CanMarkAttendance(ctx)
	if err != nil {
		return nil, err
	}

	updatedAttendance, err := uc.repo.MarkAttendance(ctx, attenance_id, scope)
	if err != nil {
		return nil, err
	}

	return updatedAttendance, nil
}
