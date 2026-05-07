package gormutils

import (
	"fmt"

	"gorm.io/gorm"

	policytypes "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
)

// apply scope in gorm_entity_repo utils. suitable to overrided READ methods.
func ApplyScope(
	db *gorm.DB,
	scope *policytypes.Scope,
	userForeignKey string,
	schoolColumn string,
) (*gorm.DB, error) {

	if scope == nil || scope.IsGlobal {
		return db, nil
	}

	// ===================== SCHOOL FILTER =====================
	if scope.School_id != "" {
		if schoolColumn == "" {
			return nil, fmt.Errorf("ApplyScope: schoolColumn is required when scope.School_id is set")
		}
		db = db.Where(fmt.Sprintf("%s = ?", schoolColumn), scope.School_id)
	}

	// ===================== USER FILTER =====================
	if scope.User_id != 0 {
		if userForeignKey == "" {
			return nil, fmt.Errorf("ApplyScope: userForeignKey is required when scope.User_id is set")
		}
		db = db.Where(fmt.Sprintf("%s = ?", userForeignKey), scope.User_id)
	}

	return db, nil
}
