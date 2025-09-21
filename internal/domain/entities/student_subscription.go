package entities

import "time"

type StudentSubscription struct {
	ID              int64
	StudentID       int64
	SubscriptionID  int64
	IssuedAt        time.Time
	ExpiresAt       time.Time
	RemainingVisits int
	IsActive        bool
	SchoolID        string
}
