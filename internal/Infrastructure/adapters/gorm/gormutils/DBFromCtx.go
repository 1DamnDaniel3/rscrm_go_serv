package gormutils

import (
	"context"

	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"gorm.io/gorm"
)

// If we have transaction in ctx - then we should use db from tx, else use db from repo
func DBFromCtx(ctx context.Context, db *gorm.DB) *gorm.DB {
	if tx, ok := ctx.Value(adapters.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return db
}
