package entities

import "time"

type SubscriptionPause struct {
	ID                    int64
	StudentSubscriptionID int64
	PausedFrom            time.Time
	PausedTo              time.Time
	SchoolID              string
}
