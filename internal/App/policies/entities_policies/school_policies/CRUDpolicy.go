package schoolpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type SchoolCrudPolicy struct{}

type ISchoolCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewEmployeeRatePolicyPolicy() ISchoolCrudPolicy {
	return &SchoolCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *SchoolCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *SchoolCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// UPDATE — только owner
func (p *SchoolCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *SchoolCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
