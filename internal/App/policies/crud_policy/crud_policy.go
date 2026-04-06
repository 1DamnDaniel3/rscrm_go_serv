package crudpolicy

import "context"

type ICRUDPolicy interface {
	CanCreate(ctx context.Context) error
	CanRead(ctx context.Context) error
	CanUpdate(ctx context.Context) error
	CanDelete(ctx context.Context) error
}
