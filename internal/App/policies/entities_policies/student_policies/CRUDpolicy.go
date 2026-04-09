package studentpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type StudentCrudPolicy struct{}

type IStudentCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewStudentCrudPolicy() IStudentCrudPolicy {
	return &StudentCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *StudentCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *StudentCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Accountant)
}

// UPDATE — только owner
func (p *StudentCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *StudentCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
