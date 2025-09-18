package entities

import "time"

type Schedule struct {
	ID              int64
	GroupID         int64
	DirectionID     int64
	TeacherID       int64
	Weekday         int
	StartTime       time.Time
	DurationMinutes int
	SchoolID        string
	ActiveFrom      time.Time
	ActiveTo        time.Time
}
