package accountrolespolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type AccountRolesCrudPolicy struct{}

type IAccountRolesCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewAccountRolesCrudPolicy() IAccountRolesCrudPolicy {
	return &AccountRolesCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *AccountRolesCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

func (p *AccountRolesCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx,
		valuetypes.Owner,
		valuetypes.Accountant,
	)
}

// UPDATE — только owner
func (p *AccountRolesCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *AccountRolesCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
