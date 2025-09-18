package entities

import "time"

type LessonSubscription struct {
	LessonID       int64
	StudentID      int64
	SubscriptionID int64
	UsedAt         time.Time
	SchoolID       string
}
