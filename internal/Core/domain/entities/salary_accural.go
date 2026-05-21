package entities

import (
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type SalaryAccural struct {
	ID              int64
	Employee_id     int64
	Period_start    time.Time
	Period_end      time.Time
	Amount          valuetypes.Money
	Status          string //CHECK (status IN ('calculated', 'approved', 'paid')),
	Calculated_at   time.Time
	TimeApproved_at time.Time
	School_id       string
}

func (l *SalaryAccural) BeforeCreate() error {
	if l.Calculated_at.IsZero() {
		l.Calculated_at = time.Now()
	}
	return nil
}
