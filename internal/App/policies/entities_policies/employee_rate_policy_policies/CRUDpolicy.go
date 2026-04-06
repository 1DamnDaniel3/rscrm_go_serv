package employeeratepolicypolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type EmployeeRatePolicyCrudPolicy struct{}

type IEmployeeRatePolicyCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewEmployeeRatePolicyPolicy() IEmployeeRatePolicyCrudPolicy {
	return &EmployeeRatePolicyCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *EmployeeRatePolicyCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *EmployeeRatePolicyCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Accountant)
}

// UPDATE — только owner
func (p *EmployeeRatePolicyCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *EmployeeRatePolicyCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
