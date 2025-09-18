package entities

import "time"

type TeacherRate struct {
	ID         int64
	TeacherID  int64
	PolicyID   int64
	ActiveFrom time.Time
	ActiveTo   time.Time
	SchoolID   string
}
