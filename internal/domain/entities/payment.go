package entities

import "time"

type Payment struct {
	ID             int64
	StudentID      int64
	SubscriptionID int64
	Amount         float64
	PaidAt         time.Time
	CreatedBy      int64
	SchoolID       string
}
