package entities

import "time"

type SubscriptionPause struct {
	ID                     int64
	StudentSubscription_id int64
	Paused_from            time.Time
	Paused_to              time.Time
	School_id              string
}
