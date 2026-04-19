package crudpolicy

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
)

type ICRUDPolicy interface {
	CanCreate(ctx context.Context) (*policytypes.Scope, error)

	CanReadOne(ctx context.Context) (*policytypes.Scope, error)
	CanReadAll(ctx context.Context) (*policytypes.Scope, error)
	CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error)

	CanUpdate(ctx context.Context) (*policytypes.Scope, error)
	CanDelete(ctx context.Context) (*policytypes.Scope, error)
}
