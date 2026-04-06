package entities

import (
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type FinancialReport struct {
	ID             int64
	Period_start   time.Time
	Period_end     time.Time
	Total_income   valuetypes.Money
	Total_expenses valuetypes.Money
	Created_at     time.Time
	School_id      string
}

func (l *FinancialReport) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
