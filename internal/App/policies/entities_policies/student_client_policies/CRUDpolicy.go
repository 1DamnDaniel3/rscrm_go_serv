package studentclientpolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type StudentClientCrudPolicy struct{}

type IStudentClientCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewStudentClientCrudPolicy() IStudentClientCrudPolicy {
	return &StudentClientCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только owner
func (p *StudentClientCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// READ — owner + accountant
func (p *StudentClientCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Accountant)
}

// UPDATE — только owner
func (p *StudentClientCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}

// DELETE — только owner
func (p *StudentClientCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner)
}
