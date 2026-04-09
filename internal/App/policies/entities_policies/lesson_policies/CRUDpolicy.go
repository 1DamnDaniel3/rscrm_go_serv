package lessonpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type LessonCrudPolicy struct{}

type ILessonCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewLessonCrudPolicy() ILessonCrudPolicy {
	return &LessonCrudPolicy{}
}

// ---====== methods ======---

func (p *LessonCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *LessonCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Teacher)
}

func (p *LessonCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *LessonCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
