package entities

import "time"

type School struct {
	ID         string
	Name       string
	City       string
	Phone      string
	Email      string
	Created_at time.Time
}
