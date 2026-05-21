package attendancepolicies

import (
	"context"
	"fmt"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type AttendanceUpdatePolicy struct{}

func NewAttendanceUpdatePolicy() IAttendanceUpdatePolicy {
	return &AttendanceUpdatePolicy{}
}

type IAttendanceUpdatePolicy interface {
	CanMarkAttendance(ctx context.Context) (*policytypes.Scope, error)
}

func (p *AttendanceUpdatePolicy) CanMarkAttendance(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	// owner + receptionist отмечать посещаемость любого занятия
	if policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Receptionist,
	) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// teacher отмечает только посещаемость на своих занятиях
	if policyutils.HasAnyRole(user,
		valuetypes.Teacher,
	) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
			User_id:   user.ID,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}
