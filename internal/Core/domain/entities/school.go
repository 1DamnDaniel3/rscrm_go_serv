package entities

import "time"

type School struct {
	Id         string
	Name       string
	City       string
	Phone      string
	Email      string
	Created_at time.Time
}
