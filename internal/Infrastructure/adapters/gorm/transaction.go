package adapters

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"gorm.io/gorm"
)

type GormTransaction struct {
	db *gorm.DB
}

func (t *GormTransaction) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		return fn(ctx)
	})
}

func NewGormTransaction(db *gorm.DB) services.Transaction {
	return &GormTransaction{db: db}
}
