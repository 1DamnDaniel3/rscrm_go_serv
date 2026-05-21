package gormentityrepos

import (
	"context"
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormUserProfileRepo struct {
	*genericAdapter.GormRepository[entities.UserProfile]
	db *gorm.DB
}

func (r *GormUserProfileRepo) Register(ctx context.Context, entity *entities.UserProfile) error {
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func NewGormUserProfileRepo(db *gorm.DB) entitiesrepos.UserProfileRepo {
	if db == nil {
		panic("NewGormUserProfileRepo: db is nil — database not initialized")
	}
	return &GormUserProfileRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.UserProfile](db),
		db:             db,
	}
}

func (r *GormUserProfileRepo) GetAllProfilesByRoles(
	ctx context.Context,
	scope *policytypes.Scope,
	roles ...string,
) (*[]entities.UserProfile, error) {

	userProfiles := &[]entities.UserProfile{}

	db := gormutils.DBFromCtx(ctx, r.db)
	db, err := gormutils.ApplyScope(db, scope, "up.id", "up.school_id")
	if err != nil {
		return nil, err
	}

	// if len(roles) > 0 {
	// 	db = db.Where(`
	// 	EXISTS (
	// 		SELECT 1
	// 		FROM account_roles ar2
	// 		JOIN roles r2 ON r2.id = ar2.role_id
	// 		WHERE ar2.account_id = up.account_id
	// 		AND r2.role IN ?
	// 	)
	// `, roles)
	// }

	if len(roles) > 0 {
		db = db.Where("r.role IN ?", roles)
	}

	err = db.Table("user_profiles up").
		Select("DISTINCT up.*").
		Joins("JOIN account_roles ar ON ar.account_id = up.account_id").
		Joins("JOIN roles r ON ar.role_id = r.id").
		Scan(userProfiles).Error
	if err != nil {
		return nil, err
	}

	return userProfiles, nil

}

// ============================ OVERRIDE ===============================

func (r *GormUserProfileRepo) GetAllWhere(
	ctx context.Context,
	filters map[string]interface{},
	entities *[]entities.UserProfile,
	scope *policytypes.Scope,
) error {

	db := gormutils.DBFromCtx(ctx, r.db)

	delete(filters, "school_id")
	delete(filters, "account_id")

	db, err := gormutils.ApplyScope(db, scope, "account_id", "school_id")
	if err != nil {
		return err
	}

	db = db.Where(filters)

	return db.Find(entities).Error
}

func (r *GormUserProfileRepo) Update(
	ctx context.Context,
	id any,
	fields map[string]interface{},
	scope *policytypes.Scope,
) error {

	delete(fields, "school_id")
	delete(fields, "id")

	db := gormutils.DBFromCtx(ctx, r.db)
	db, err := gormutils.ApplyScope(db, scope, "account_id", "school_id")
	if err != nil {
		return err
	}

	tx := db.
		Model(&entities.UserProfile{}).
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
