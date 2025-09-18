package entities

import "time"

type School struct {
	ID        string
	Name      string
	City      string
	Phone     string
	Email     string
	CreatedAt time.Time
}
