package entities

import "time"

type Lesson struct {
	ID              int64
	GroupID         int64
	DirectionID     int64
	TeacherID       int64
	LessonDate      time.Time
	StartTime       time.Time
	DurationMinutes int
	IsCanceled      bool
	SchoolID        string
}
