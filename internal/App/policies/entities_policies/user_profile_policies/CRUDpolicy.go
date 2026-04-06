package userprofilepolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type UserProfieCrudPolicy struct{}

type IUserProfieCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewUserProfieCrudPolicy() IUserProfieCrudPolicy {
	return &UserProfieCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *UserProfieCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + Manager
func (p *UserProfieCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// UPDATE — только owner
func (p *UserProfieCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *UserProfieCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
