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
	foreignKey string, // например: "teacher_id", "created_by", "account_id"
) *gorm.DB {

	if scope == nil || scope.IsGlobal {
		return db
	}

	// ===================== SCHOOL FILTER =====================
	if scope.School_id != "" {
		db = db.Where("school_id = ?", scope.School_id)
	}

	// ===================== USER FILTER =====================
	if scope.User_id != 0 && foreignKey != "" {
		db = db.Where(fmt.Sprintf("%s = ?", foreignKey), scope.User_id)
	}

	return db
}
