package entities

import "time"

type Subscription struct {
	ID          int64
	Name        string
	Price       float64
	Visit_limit int
	Active_from time.Time
	Active_to   time.Time
	Is_archived bool
	School_id   string
}
