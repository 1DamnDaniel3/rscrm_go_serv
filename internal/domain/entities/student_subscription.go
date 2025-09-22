package entities

import "time"

type StudentSubscription struct {
	ID               int64
	Student_id       int64
	Subscription_id  int64
	Issued_at        time.Time
	Expires_at       time.Time
	Remaining_visits int
	Is_active        bool
	School_id        string
}
