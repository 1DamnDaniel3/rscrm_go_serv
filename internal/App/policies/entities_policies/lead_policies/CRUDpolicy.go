package leadpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type LeadCrudPolicy struct{}

type ILeadCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewLeadCrudPolicy() ILeadCrudPolicy {
	return &LeadCrudPolicy{}
}

// ---====== methods ======---

func (p *LeadCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Receptionist)
}

func (p *LeadCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Receptionist)
}

func (p *LeadCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

func (p *LeadCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
