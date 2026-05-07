package lessonpolicies

import (
	"context"
	"fmt"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type LessonCreatePolicy struct{}

type ILessonCreatePolicy interface {
	CanGenerateLessons(ctx context.Context) (*policytypes.Scope, error)
}

func NewLessonCreatePolicy() ILessonCreatePolicy {
	return &LessonCreatePolicy{}
}

func (p *LessonCreatePolicy) CanGenerateLessons(ctx context.Context) (*policytypes.Scope, error) {

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Teacher,
		valuetypes.Receptionist,
	) {
		return nil, fmt.Errorf("forbidden.")
	}

	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil

}
