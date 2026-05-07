package schedulepolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type ScheduleCrudPolicy struct{}

type IScheduleCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewScheduleCrudPolicy() IScheduleCrudPolicy {
	return &ScheduleCrudPolicy{}
}

// ---====== methods ======---

// CREATE — Owner + Manager
func (p *ScheduleCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
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
func (p *ScheduleCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE — Owner + Manager + Teacher
func (p *ScheduleCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Teacher,
		valuetypes.Receptionist,
	) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}

// READ ALL
func (p *ScheduleCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *ScheduleCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE — Owner + Manager
func (p *ScheduleCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
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

// DELETE — только Owner
func (p *ScheduleCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Owner) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil
}
