package leadgrouppolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type LeadGroupCrudPolicy struct{}

type ILeadGroupCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewEmployeeRatePolicyPolicy() ILeadGroupCrudPolicy {
	return &LeadGroupCrudPolicy{}
}

// ---====== methods ======---

func (p *LeadGroupCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *LeadGroupCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Receptionist)
}

func (p *LeadGroupCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *LeadGroupCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}
