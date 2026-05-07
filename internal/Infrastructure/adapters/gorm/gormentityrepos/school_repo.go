package gormentityrepos

import (
	"context"
	"errors"
	"fmt"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormSchoolRepo struct {
	*genericAdapter.GormRepository[entities.School]
	db *gorm.DB
}

func NewGormSchoolRepo(db *gorm.DB) entitiesrepos.SchoolRepository {
	return &GormSchoolRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.School](db),
		db:             db,
	}
}

func (r *GormSchoolRepo) Register(ctx context.Context, entity *entities.School) error {
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

// ============================ OVERRIDE ============================

// ==================================================== GetByID
func (r *GormSchoolRepo) GetByID(
	ctx context.Context,
	entityID any,
	entity *entities.School,
	scope *policytypes.Scope,
) error {

	db := gormutils.DBFromCtx(ctx, r.db)

	// ===================== ACCESS CONTROL =====================
	if scope != nil && !scope.IsGlobal {
		if scope.School_id == "" {
			return fmt.Errorf("forbidden")
		}

		// Фильтр
		db = db.Where("id = ?", scope.School_id)
	}

	return db.First(entity, "id = ?", entityID).Error
}

// ==================================================== GetAllWhere
func (r *GormSchoolRepo) GetAllWhere(
	ctx context.Context,
	filters map[string]interface{},
	entities *[]entities.School,
	scope *policytypes.Scope,
) error {
	db := gormutils.DBFromCtx(ctx, r.db)

	// ===================== ACCESS CONTROL =====================
	if scope != nil && !scope.IsGlobal {
		if scope.School_id == "" {
			return fmt.Errorf("forbidden")
		}

		// Фильтр
		db = db.Where("id = ?", scope.School_id)
	}

	return db.Where(filters).Find(entities).Error
}

// ==================================================== GetAll
func (r *GormSchoolRepo) GetAll(
	ctx context.Context,
	entities *[]entities.School,
	scope *policytypes.Scope,
) error {
	db := gormutils.DBFromCtx(ctx, r.db)

	// ===================== ACCESS CONTROL =====================
	if scope != nil && !scope.IsGlobal {
		if scope.School_id == "" {
			return fmt.Errorf("forbidden")
		}

		// Фильтр
		db = db.Where("id = ?", scope.School_id)
	}

	// Фильтруем только свою школу
	return db.Find(entities).Error
}

// ==================================================== UPDATE
func (r *GormSchoolRepo) Update(
	ctx context.Context,
	entityID any,
	fields map[string]interface{},
	scope *policytypes.Scope,
) error {

	if len(fields) == 0 {
		return nil
	}

	delete(fields, "id")
	delete(fields, "created_at")

	db := gormutils.DBFromCtx(ctx, r.db)

	// ===================== ACCESS CONTROL =====================
	if scope != nil && !scope.IsGlobal {

		if scope.School_id == "" {
			return errors.New("forbidden")
		}

		db = db.Where("id = ?", scope.School_id)
	}

	// ===================== UPDATE =====================
	tx := db.
		Model(&entities.School{}).
		Where("id = ?", entityID).
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
func (r *GormSchoolRepo) Delete(
	ctx context.Context,
	entityID any,
	entity *entities.School,
	scope *policytypes.Scope,
) error {
	db := gormutils.DBFromCtx(ctx, r.db)

	if scope == nil || (!scope.IsGlobal && scope.School_id == "") {
		return errors.New("invalid scope")
	}

	tx := db.
		Where("id = ?", entityID)

	// Если не global — ограничиваем своей школой
	if !scope.IsGlobal {
		tx = tx.Where("id = ?", scope.School_id)
	}

	tx = tx.Delete(entity)

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("entity not found or access denied")
	}

	return nil
}
