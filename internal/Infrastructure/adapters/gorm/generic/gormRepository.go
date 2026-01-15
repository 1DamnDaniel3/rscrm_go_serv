package generic

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

// If we have transaction in ctx - then we should use db from tx, else use db from repo
func (r *GormRepository[T]) dbFromCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return r.db
}

func (r *GormRepository[T]) Create(ctx context.Context, entity *T) error {

	if beforeCreate, ok := any(entity).(services.BeforeCreate); ok {
		if err := beforeCreate.BeforeCreate(); err != nil {
			return err
		}
	}

	return r.dbFromCtx(ctx).Create(entity).Error
}

func (r *GormRepository[T]) GetByID(id any, entity *T) error {
	return r.db.First(entity, "id=?", id).Error
}

func (r *GormRepository[T]) GetAllWhere(filters map[string]interface{}, entities *[]T) error {
	return r.db.Where(filters).Find(entities).Error
}

func (r *GormRepository[T]) GetAll(entities *[]T) error {
	return r.db.Find(entities).Error
}

func (r *GormRepository[T]) Update(id any, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}
	return r.db.Model(new(T)).Where("id = ?", id).Updates(fields).Error
}

func (r *GormRepository[T]) Delete(id any, entity *T) error {
	return r.db.Delete(entity, "id=?", id).Error
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}
