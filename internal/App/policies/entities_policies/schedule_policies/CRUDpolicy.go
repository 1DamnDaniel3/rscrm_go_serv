package schedulepolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type ScheduleCrudPolicy struct{}

type IScheduleCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewEmployeeRatePolicyPolicy() IScheduleCrudPolicy {
	return &ScheduleCrudPolicy{}
}

// ---====== methods ======---

func (p *ScheduleCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *ScheduleCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Teacher)
}

func (p *ScheduleCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *ScheduleCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
