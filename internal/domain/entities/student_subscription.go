package entities

import "time"

type StudentSubscription struct {
	StudentID       int64
	SubscriptionID  int64
	IssuedAt        time.Time
	ExpiresAt       time.Time
	RemainingVisits int
	IsActive        bool
	SchoolID        string
}
