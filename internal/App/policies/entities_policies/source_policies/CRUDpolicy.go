package sourcepolicies

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
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

// CREATE — только owner
func (p *SourceCrudPolicy) CanCreate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// READ — owner + accountant
func (p *SourceCrudPolicy) CanRead(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// UPDATE — только owner
func (p *SourceCrudPolicy) CanUpdate(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}

// DELETE — только owner
func (p *SourceCrudPolicy) CanDelete(ctx context.Context) error {
	return policyutils.RequireRoles(ctx, valuetypes.Owner, valuetypes.Manager)
}
