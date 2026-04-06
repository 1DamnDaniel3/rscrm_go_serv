package useraccountpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type UserAccountCrudPolicy struct{}

type IUserAccountCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewUserAccountCrudPolicy() IUserAccountCrudPolicy {
	return &UserAccountCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *UserAccountCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *UserAccountCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// UPDATE — только owner
func (p *UserAccountCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *UserAccountCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
