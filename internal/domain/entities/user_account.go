package entities

import "time"

type UserAccount struct {
	ID         int64
	Email      string
	Password   string
	Role       string
	Created_at time.Time
	School_id  string
}
