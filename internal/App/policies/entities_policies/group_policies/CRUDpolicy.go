package grouppolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
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

// CREATE — Owner + Manager (в рамках своей школы)
func (p *GroupCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Owner, valuetypes.Manager) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil
}

// CREATE MANY
func (p *GroupCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE
func (p *GroupCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Owner + Manager + Accountant внутри школы
	if policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Teacher,
		valuetypes.Accountant,
	) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}

// READ ALL
func (p *GroupCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *GroupCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE — Owner + Manager
func (p *GroupCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Owner, valuetypes.Manager) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil
}

// DELETE — Owner + Manager
func (p *GroupCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Owner, valuetypes.Manager) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil
}
