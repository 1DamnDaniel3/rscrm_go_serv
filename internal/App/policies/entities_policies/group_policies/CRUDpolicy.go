package grouppolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type GroupCrudPolicy struct{}

type IGroupCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewGroupCrudPolicy() IGroupCrudPolicy {
	return &GroupCrudPolicy{}
}

// ---====== methods ======---

// CREATE — owner + manager
func (p *GroupCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// READ — owner + manager + accountant
func (p *GroupCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager, valuetypes.Accountant)
}

// UPDATE — owner + manager
func (p *GroupCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// DELETE — owner + manager
func (p *GroupCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}
