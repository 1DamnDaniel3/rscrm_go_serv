package subscriptionpolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type SubscriptionCrudPolicy struct{}

type ISubscriptionCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewSubscriptionCrudPolicy() ISubscriptionCrudPolicy {
	return &SubscriptionCrudPolicy{}
}

// ---====== methods ======---

// CREATE — только Owner
func (p *SubscriptionCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
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

// CREATE MANY
func (p *SubscriptionCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE — Owner + Accountant
func (p *SubscriptionCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Teacher,
		valuetypes.Accountant,
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
func (p *SubscriptionCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *SubscriptionCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE — только Owner
func (p *SubscriptionCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
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

// DELETE — только Owner
func (p *SubscriptionCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
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
