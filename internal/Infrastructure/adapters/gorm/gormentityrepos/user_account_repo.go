package gormentityrepos

import (
	"context"
	"errors"
	"strings"
	"time"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormUserAccountRepo struct {
	*genericAdapter.GormRepository[entities.UserAccount]
	db     *gorm.DB
	hasher ports.PasswordHasher
}

func NewGormUserAccountRepo(db *gorm.DB, hasher ports.PasswordHasher) entitiesrepos.UserAccountRepository {
	return &GormUserAccountRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.UserAccount](db),
		db:             db,
		hasher:         hasher}
}

func (r *GormUserAccountRepo) Register(ctx context.Context, entity *entities.UserAccount) error {
	var err error
	entity.Password, err = r.hasher.Hash(entity.Password)
	if err != nil {
		return err
	}
	tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB)
	if !ok {
		return errors.New("no transaction in context")
	}
	return tx.Create(entity).Error
}

func (r *GormUserAccountRepo) GetByEmail(email string) (*entities.UserAccount, error) {
	var user entities.UserAccount
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ======================================== OVERRIDE =============================================

func (r *GormUserAccountRepo) Create(
	ctx context.Context,
	entity *entities.UserAccount,
	scope *policytypes.Scope,
) error {

	db := gormutils.DBFromCtx(ctx, r.db)

	// scope
	if scope != nil && !scope.IsGlobal {
		entity.School_id = scope.School_id
	}

	var err error
	entity.Password, err = r.hasher.Hash(entity.Password)
	if err != nil {
		return err
	}

	if beforeCreate, ok := any(entity).(services.BeforeCreate); ok {
		if err := beforeCreate.BeforeCreate(); err != nil {
			return err
		}
	}

	return db.Create(entity).Error
}

// -=== - === - === - === - === - === UPDATE

func (r *GormUserAccountRepo) Update(
	ctx context.Context,
	id any,
	fields map[string]interface{},
	scope *policytypes.Scope,
) error {

	delete(fields, "school_id") // can't change
	delete(fields, "account_id")
	delete(fields, "id")

	db := gormutils.DBFromCtx(ctx, r.db)
	db, err := gormutils.ApplyScope(db, scope, "id", "school_id")
	if err != nil {
		return err
	}

	// if password is changing - hash it
	if pwd, ok := fields["password"].(string); ok && pwd != "" {
		hashed, err := r.hasher.Hash(pwd)
		if err != nil {
			return err
		}
		fields["password"] = hashed
	}

	tx := db.
		Model(&entities.UserAccount{}).
		Where("id = ?", id).
		Updates(fields)

	delete(fields, "password") // Хеш не отправляем обратно

	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("entity not found or access denied")
	}

	return nil
}

// -=== - === - === - === - === - === GetByID

func (r *GormUserAccountRepo) GetByID(
	ctx context.Context,
	id any,
	entity *entities.UserAccount,
	scope *policytypes.Scope,
) error {

	db := r.DBFromCtx(ctx)
	db, err := gormutils.ApplyScope(db, scope, "id", "school_id")
	if err != nil {
		return err
	}

	return db.First(entity, "id = ?", id).Error
}

// ======================================== QUERY SERVICE =============================================

type GormUserAccountQueryService struct {
	db *gorm.DB
}

func NewGormUserAccountQueryService(db *gorm.DB) entitiesrepos.UserAccountQueryService {
	return &GormUserAccountQueryService{db}
}

func (r *GormUserAccountQueryService) GetMe(ctx context.Context) (*businessobjects.UserBO, error) {
	db := gormutils.DBFromCtx(ctx, r.db)

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	type result struct {
		ID        int64
		Email     string
		Password  string
		CreatedAt time.Time
		SchoolID  string
		RolesRaw  string
	}

	var res result

	err = db.Table("user_accounts ua").
		Select(`
			ua.id,
			ua.email,
			ua.password,
			ua.created_at,
			ua.school_id,
			COALESCE(array_to_string(array_agg(r.role), ','), '') as roles_raw
		`).
		Joins("LEFT JOIN account_roles ar ON ar.account_id = ua.id").
		Joins("LEFT JOIN roles r ON r.id = ar.role_id").
		Where("ua.id = ?", user.ID).
		Group("ua.id").
		Scan(&res).Error

	if err != nil {
		return nil, err
	}

	if res.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	var roles []string
	if res.RolesRaw != "" {
		roles = strings.Split(res.RolesRaw, ",")
	}

	userBO := &businessobjects.UserBO{
		UserAccount: entities.UserAccount{
			Id:         res.ID,
			Email:      res.Email,
			Password:   res.Password,
			Created_at: res.CreatedAt,
			School_id:  res.SchoolID,
		},
		Roles: roles,
	}

	return userBO, nil
}

func (r *GormUserAccountQueryService) GetAllAccountsWithRoles(
	ctx context.Context,
	scope *policytypes.Scope,
) ([]*businessobjects.UserBO, error) {

	db := gormutils.DBFromCtx(ctx, r.db)
	db, err := gormutils.ApplyScope(db, scope, "ua.id", "ua.school_id")
	if err != nil {
		return nil, err
	}

	type result struct {
		ID        int64
		Email     string
		Password  string
		CreatedAt time.Time
		SchoolID  string
		RolesRaw  string
	}

	var results []result

	err = db.Table("user_accounts ua").
		Select(`
            ua.id,
            ua.email,
            ua.password,
            ua.created_at,
            ua.school_id,
            COALESCE(array_to_string(array_agg(r.role), ','), '') as roles_raw
        `).
		Joins("LEFT JOIN account_roles ar ON ar.account_id = ua.id").
		Joins("LEFT JOIN roles r ON r.id = ar.role_id").
		Group("ua.id").
		Scan(&results).Error

	if err != nil {
		return nil, err
	}

	var users []*businessobjects.UserBO

	for _, res := range results {
		var roles []string
		if res.RolesRaw != "" {
			roles = strings.Split(res.RolesRaw, ",")
		}

		userBO := &businessobjects.UserBO{
			UserAccount: entities.UserAccount{
				Id:         res.ID,
				Email:      res.Email,
				Password:   res.Password,
				Created_at: res.CreatedAt,
				School_id:  res.SchoolID,
			},
			Roles: roles,
		}

		users = append(users, userBO)
	}

	return users, nil
}
