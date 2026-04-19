package genericrepo

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
)

type Repository[T any] interface {
	Create(ctx context.Context, entity *T, scope *policytypes.Scope) error
	CreateMany(ctx context.Context, entities *[]T, scope *policytypes.Scope) error
	GetByID(ctx context.Context, id any, entity *T, scope *policytypes.Scope) error
	GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T, scope *policytypes.Scope) error
	GetAll(ctx context.Context, entities *[]T, scope *policytypes.Scope) error
	Update(ctx context.Context, id any, fields map[string]interface{}, scope *policytypes.Scope) error
	Delete(ctx context.Context, id any, entity *T, scope *policytypes.Scope) error

	FindRelation(ctx context.Context, relationMap map[string]any, scope *policytypes.Scope) (*T, error)
}
