package entities

import "time"

type FinancialReport struct {
	ID             int64
	Period_start   time.Time
	Period_end     time.Time
	Total_income   float64
	Total_expenses float64
	Created_at     time.Time
	School_id      string
}

func (l *FinancialReport) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
