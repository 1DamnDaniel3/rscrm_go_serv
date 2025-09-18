package entities

import "time"

type Attendance struct {
	StudentID int64
	LessonID  int64
	Status    string
	MarkedBy  int64
	MarkedAt  time.Time
	SchoolID  string
}
