package useraccountpolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
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

// CREATE — только Owner (создаёт аккаунты в своей школе)
func (p *UserAccountCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
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

// CREATE MANY — та же логика
func (p *UserAccountCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE
func (p *UserAccountCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → все аккаунты
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → все аккаунты своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// Остальные → только свой аккаунт
	return &policytypes.Scope{
		IsGlobal: false,
		User_id:  user.ID,
	}, nil
}

// READ ALL
func (p *UserAccountCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *UserAccountCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE
func (p *UserAccountCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → любые аккаунты
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → аккаунты своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// Остальные → только свой аккаунт
	return &policytypes.Scope{
		IsGlobal: false,
		User_id:  user.ID,
	}, nil
}

// DELETE
func (p *UserAccountCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → любые аккаунты
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → аккаунты своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}
