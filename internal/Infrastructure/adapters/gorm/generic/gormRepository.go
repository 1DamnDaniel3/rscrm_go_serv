package generic

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

// PROTECT ONLY SCHOOL ACCESS with scopes from policy layer
// if you need to protect user by user_id OVERRIDE this methods in his gorm_entity_repos
func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}

// ==================================================== Create
func (r *GormRepository[T]) Create(ctx context.Context, entity *T, scope *policytypes.Scope) error {
	db := r.DBFromCtx(ctx)

	// ===================== APPLY SCHOOL SCOPE =====================
	if scope != nil && !scope.IsGlobal && scope.School_id != "" {

		val := reflect.ValueOf(entity).Elem()

		field := val.FieldByName("School_id")
		if field.IsValid() && field.CanSet() {

			// если уже есть значение - проверяем
			if field.String() != "" && field.String() != scope.School_id {
				return fmt.Errorf("forbidden: invalid school_id")
			}

			field.SetString(scope.School_id)
		}
	}

	// ===================== HOOK =====================
	if beforeCreate, ok := any(entity).(services.BeforeCreate); ok {
		if err := beforeCreate.BeforeCreate(); err != nil {
			return err
		}
	}

	// ===================== CREATE =====================
	return db.Create(entity).Error
}

// ==================================================== CreateMany
func (r *GormRepository[T]) CreateMany(ctx context.Context, entities *[]T, scope *policytypes.Scope) error {
	db := r.DBFromCtx(ctx)

	// ===================== APPLY SCOPE =====================
	if scope != nil && !scope.IsGlobal && scope.School_id != "" {

		val := reflect.ValueOf(entities).Elem()

		for i := 0; i < val.Len(); i++ {
			entityVal := val.Index(i)

			field := entityVal.FieldByName("School_id")
			if field.IsValid() && field.CanSet() {

				// защита от подмены
				if field.String() != "" && field.String() != scope.School_id {
					return fmt.Errorf("forbidden: invalid school_id")
				}

				field.SetString(scope.School_id)
			}

			// ===================== HOOK =====================
			entityPtr := entityVal.Addr().Interface()
			if beforeCreate, ok := entityPtr.(services.BeforeCreate); ok {
				if err := beforeCreate.BeforeCreate(); err != nil {
					return err
				}
			}
		}
	}

	// ===================== INSERT =====================
	return db.
		Clauses(clause.OnConflict{
			DoNothing: true,
		}).
		CreateInBatches(entities, 500).
		Error
}

// ==================================================== GetByID

func (r *GormRepository[T]) GetByID(ctx context.Context, id any, entity *T, scope *policytypes.Scope) error {

	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

	return db.First(entity, "id = ?", id).Error
}

// ==================================================== GetAllWhere

func (r *GormRepository[T]) GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T, scope *policytypes.Scope) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

	delete(filters, "school_id")

	return db.Where(filters).Find(entities).Error
}

// ==================================================== GetAll

func (r *GormRepository[T]) GetAll(ctx context.Context, entities *[]T, scope *policytypes.Scope) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

	return db.Find(entities).Error
}

// ==================================================== Update

func (r *GormRepository[T]) Update(ctx context.Context, id any, fields map[string]interface{}, scope *policytypes.Scope) error {
	if len(fields) == 0 {
		return nil
	}

	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

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

func (r *GormRepository[T]) Delete(ctx context.Context, id any, entity *T, scope *policytypes.Scope) error {
	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

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
func (r *GormRepository[T]) FindRelation(ctx context.Context, relationMap map[string]any, scope *policytypes.Scope) (*T, error) {
	db := r.DBFromCtx(ctx)
	db = r.ApplyScope(db, scope)

	var entity T
	err := db.Where(relationMap).First(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// --- ========================================== utils ========================================== ---

func (r *GormRepository[T]) ApplyScope(
	db *gorm.DB,
	scope *policytypes.Scope,
) *gorm.DB {

	if scope == nil || scope.IsGlobal {
		return db
	}

	// Проверяем есть ли поле School_id у модели
	if _, ok := reflect.TypeOf(new(T)).Elem().FieldByName("School_id"); ok {
		if scope.School_id != "" {
			db = db.Where("school_id = ?", scope.School_id)
		}
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
