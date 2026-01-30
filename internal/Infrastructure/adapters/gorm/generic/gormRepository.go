package generic

import (
	"context"
	"errors"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func (r *GormRepository[T]) Create(ctx context.Context, entity *T) error {

	if beforeCreate, ok := any(entity).(services.BeforeCreate); ok {
		if err := beforeCreate.BeforeCreate(); err != nil {
			return err
		}
	}

	return r.dbFromCtx(ctx).Create(entity).Error
}

func (r *GormRepository[T]) GetByID(ctx context.Context, id any, entity *T) error {
	db := r.dbFromCtx(ctx)
	db = r.applyTenantFilter(ctx, db)

	return db.First(entity, "id = ?", id).Error
}

func (r *GormRepository[T]) GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T) error {
	db := r.dbFromCtx(ctx)
	db = r.applyTenantFilter(ctx, db)

	delete(filters, "school_id")

	return db.Where(filters).Find(entities).Error
}

func (r *GormRepository[T]) GetAll(ctx context.Context, entities *[]T) error {
	db := r.dbFromCtx(ctx)
	db = r.applyTenantFilter(ctx, db)

	return db.Find(entities).Error
}

func (r *GormRepository[T]) Update(ctx context.Context, id any, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	db := r.dbFromCtx(ctx)
	db = r.applyTenantFilter(ctx, db)

	delete(fields, "school_id")

	tx := db.
		Model(new(T)).
		Where("id = ?", id).
		Updates(fields)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("entity not found or access denied")
	}

	return nil
}

func (r *GormRepository[T]) Delete(ctx context.Context, id any, entity *T) error {
	db := r.dbFromCtx(ctx)
	db = r.applyTenantFilter(ctx, db)

	tx := db.
		Model(entity).
		Where("id = ?", id).
		Delete(entity)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("entity not found or access denied")
	}

	return nil
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

// --- ============== utils ============== ---

func (r *GormRepository[T]) applyTenantFilter(
	ctx context.Context,
	db *gorm.DB,
) *gorm.DB {

	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return db
	}

	// Проверяем: есть ли поле School_id у модели
	if _, ok := reflect.TypeOf(new(T)).Elem().FieldByName("School_id"); ok {
		return db.Where("school_id = ?", schoolID)
	}

	return db
}

// If we have transaction in ctx - then we should use db from tx, else use db from repo
func (r *GormRepository[T]) dbFromCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return r.db
}
