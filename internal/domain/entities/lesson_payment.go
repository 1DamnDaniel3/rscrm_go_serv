package entities

import "time"

type LessonPayment struct {
	LessonID  int64
	StudentID int64
	Amount    float64
	PaidAt    time.Time
	SchoolID  string
}
