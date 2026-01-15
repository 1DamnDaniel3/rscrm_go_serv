package entities

import "time"

type TeacherRatePolicy struct {
	ID          int64
	Name        string
	Description string
	Created_at  time.Time
	School_id   string
}

func (l *TeacherRatePolicy) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
