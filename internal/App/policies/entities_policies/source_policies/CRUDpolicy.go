package sourcepolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type SourceCrudPolicy struct{}

type ISourceCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewSourceCrudPolicy() ISourceCrudPolicy {
	return &SourceCrudPolicy{}
}

// ---====== methods ======---

// CREATE — запрещено всем (кроме Admin)
func (p *SourceCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
	user, _ := policyutils.GetUserFromCtx(ctx)

	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{IsGlobal: true}, nil
	}

	return nil, fmt.Errorf("forbidden")
}

// CREATE MANY — то же самое
func (p *SourceCrudPolicy) CanCreateMany(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanCreate(ctx)
}

// READ ONE — Owner + Manager + Admin
func (p *SourceCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if policyutils.HasAnyRole(user,
		valuetypes.Owner,
		valuetypes.Manager,
		valuetypes.Admin,
	) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	return nil, fmt.Errorf("forbidden")
}

// READ ALL
func (p *SourceCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// READ ALL WHERE
func (p *SourceCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// UPDATE — запрещено всем (кроме Admin)
func (p *SourceCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
	user, _ := policyutils.GetUserFromCtx(ctx)

	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{IsGlobal: true}, nil
	}

	return nil, fmt.Errorf("forbidden")
}

// DELETE — запрещено всем (кроме Admin)
func (p *SourceCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, _ := policyutils.GetUserFromCtx(ctx)

	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{IsGlobal: true}, nil
	}

	return nil, fmt.Errorf("forbidden")
}
