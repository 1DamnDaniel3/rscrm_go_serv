package entities

import "time"

type TeacherRatePolicy struct {
	ID          int64
	Name        string
	Description string
	Created_at  time.Time
	School_id   string
}
