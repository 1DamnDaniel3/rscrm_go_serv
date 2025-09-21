package entities

import "time"

type LessonSubscription struct {
	ID             int64
	LessonID       int64
	StudentID      int64
	SubscriptionID int64
	UsedAt         time.Time
	SchoolID       string
}
