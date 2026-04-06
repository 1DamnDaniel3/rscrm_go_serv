package entities

import (
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type Transaction struct {
	Id             int64
	Type           string // CHECK (type IN ('income', 'expense')),
	Category       string // lesson or subscription_payment or rent
	Amount         valuetypes.Money
	Reference_type string // lesson/subscription
	Reference_id   int64
	Created_at     time.Time
	Created_by     int64
	School_id      string
}

func (l *Transaction) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
