package generic

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

// ==================================================== Create
func (r *GormRepository[T]) Create(ctx context.Context, entity *T) error {

	db := r.DBFromCtx(ctx)
	// Получаем school_id из контекста
	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	// Через reflect подставляем school_id, если такое поле есть
	val := reflect.ValueOf(entity).Elem()
	if field := val.FieldByName("School_id"); field.IsValid() && field.CanSet() {
		field.SetString(schoolID)
	}

	if beforeCreate, ok := any(entity).(services.BeforeCreate); ok {
		if err := beforeCreate.BeforeCreate(); err != nil {
			return err
		}
	}

	return db.Create(entity).Error
}

// ==================================================== CreateMany
func (r *GormRepository[T]) CreateMany(ctx context.Context, entities *[]T) error {
	db := r.DBFromCtx(ctx)

	// school_id из контекста
	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	val := reflect.ValueOf(entities).Elem()

	for i := 0; i < val.Len(); i++ {
		entityVal := val.Index(i)

		// Устанавливаем school_id если есть поле
		if field := entityVal.FieldByName("School_id"); field.IsValid() && field.CanSet() {
			field.SetString(schoolID)
		}

		// BeforeCreate хук
		entityPtr := entityVal.Addr().Interface()
		if beforeCreate, ok := entityPtr.(services.BeforeCreate); ok {
			if err := beforeCreate.BeforeCreate(); err != nil {
				return err
			}
		}
	}

	return db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		CreateInBatches(entities, 500).
		Error
}

// ==================================================== GetById
func (r *GormRepository[T]) GetByID(ctx context.Context, id any, entity *T) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

	return db.First(entity, "id = ?", id).Error
}

// ==================================================== GetAllWhere

func (r *GormRepository[T]) GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

	delete(filters, "school_id")

	return db.Where(filters).Find(entities).Error
}

// ==================================================== GetAll

func (r *GormRepository[T]) GetAll(ctx context.Context, entities *[]T) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

	return db.Find(entities).Error
}

// ==================================================== Update

func (r *GormRepository[T]) Update(ctx context.Context, id any, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}

	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

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

// ==================================================== Delete

func (r *GormRepository[T]) Delete(ctx context.Context, id any, entity *T) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

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

// ==================================================== FindRelation

// FindRelation by map where keys is column name(string) value is ID (int64 usually)
// expample: relationMap := map[string]any{"user_id": user.ID, "school_id: user.School_ID"}
func (r *GormRepository[T]) FindRelation(ctx context.Context, relationMap map[string]any) (*T, error) {
	db := r.DBFromCtx(ctx)
	db = r.ApplyTenantFilter(ctx, db)

	var entity T
	err := db.Where(relationMap).First(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// --- ========================================== utils ========================================== ---

func (r *GormRepository[T]) ApplyTenantFilter(
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
func (r *GormRepository[T]) DBFromCtx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return r.db
}
