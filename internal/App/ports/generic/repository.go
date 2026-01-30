package generic

import "context"

type Repository[T any] interface {
	Create(ctx context.Context, entity *T) error
	GetByID(ctx context.Context, id any, entity *T) error
	GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T) error
	GetAll(ctx context.Context, entities *[]T) error
	Update(ctx context.Context, id any, fields map[string]interface{}) error
	Delete(ctx context.Context, id any, entity *T) error
}
