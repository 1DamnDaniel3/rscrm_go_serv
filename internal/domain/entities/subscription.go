package entities

import "time"

type Subscription struct {
	ID         int64
	Name       string
	Price      float64
	VisitLimit int
	ActiveFrom time.Time
	ActiveTo   time.Time
	IsArchived bool
	SchoolID   string
}
