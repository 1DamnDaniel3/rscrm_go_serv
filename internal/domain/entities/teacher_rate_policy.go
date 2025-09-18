package entities

import "time"

type TeacherRatePolicy struct {
	ID          int64
	Name        string
	Description string
	CreatedAt   time.Time
	SchoolID    string
}
