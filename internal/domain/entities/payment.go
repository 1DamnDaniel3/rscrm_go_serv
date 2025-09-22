package entities

import "time"

type Payment struct {
	ID              int64
	Student_id      int64
	Subscription_id int64
	Amount          float64
	Paid_at         time.Time
	Created_by      int64
	School_id       string
}
