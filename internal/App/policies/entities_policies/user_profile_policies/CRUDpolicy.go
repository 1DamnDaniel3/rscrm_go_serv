package userprofilepolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
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

// CREATE — Owner создаёт профили в своей школе
func (p *UserProfieCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
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
func (p *UserProfieCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE
func (p *UserProfieCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → все профили
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner / Manager → все профили своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner, valuetypes.Manager) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// Остальные → только свой профиль (через account_id)
	return &policytypes.Scope{
		IsGlobal: false,
		User_id:  user.ID,
	}, nil
}

// READ ALL
func (p *UserProfieCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *UserProfieCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE
func (p *UserProfieCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → любые профили
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → профили своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// Остальные → только свой профиль
	return &policytypes.Scope{
		IsGlobal: false,
		User_id:  user.ID,
	}, nil
}

// DELETE
func (p *UserProfieCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → любые профили
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → профили своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}
