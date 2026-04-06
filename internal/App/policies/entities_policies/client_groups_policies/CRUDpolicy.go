package clientgroupspolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type ClientGroupsCrudPolicy struct{}

type IClientGroupsCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewClientGroupsCrudPolicy() IClientGroupsCrudPolicy {
	return &ClientGroupsCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *ClientGroupsCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — ALL ROLES
func (p *ClientGroupsCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Teacher,
		valuetypes.Receptionist,
		valuetypes.Accountant,
	)
}

// UPDATE — только owner
func (p *ClientGroupsCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *ClientGroupsCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
