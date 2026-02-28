package gormutils

import (
	"context"
	"reflect"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	"gorm.io/gorm"
)

func ApplyTenantFilter[T any](
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
