package entities

import "time"

type FinancialReport struct {
	ID            int64
	PeriodStart   time.Time
	PeriodEnd     time.Time
	TotalIncome   float64
	TotalExpenses float64
	CreatedAt     time.Time
	SchoolID      string
}
