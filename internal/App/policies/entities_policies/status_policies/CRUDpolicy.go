package statuspolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type StatusCrudPolicy struct{}

type IStatusCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewEmployeeRatePolicyPolicy() IStatusCrudPolicy {
	return &StatusCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *StatusCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *StatusCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Accountant)
}

// UPDATE — только owner
func (p *StatusCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *StatusCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
