package gormutils

import (
	"context"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	"gorm.io/gorm"
)

func ApplyTenantFilter[T any](
	ctx context.Context,
	db *gorm.DB,
) *gorm.DB {

	userCtx, ok := ctx.Value(contextkeys.User).(*valuetypes.UserContext)
	if !ok || userCtx.SchoolID == "" {
		// Если контекст пустой или нет school_id, не фильтруем
		return db
	}

	// Проверяем: есть ли поле School_id у модели
	if _, ok := reflect.TypeOf(new(T)).Elem().FieldByName("School_id"); ok {
		return db.Where("school_id = ?", userCtx.SchoolID)
	}

	return db
}
