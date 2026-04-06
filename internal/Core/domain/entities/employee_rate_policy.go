package entities

import "time"

type EmployeeRatePolicy struct {
	ID          int64
	Name        string
	Description string
	Created_at  time.Time
	School_id   string
}

func (l *EmployeeRatePolicy) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
