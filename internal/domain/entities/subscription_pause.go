package entities

import "time"

type SubscriptionPause struct {
	StudentSubscriptionID int64
	PausedFrom            time.Time
	PausedTo              time.Time
	SchoolID              string
}
