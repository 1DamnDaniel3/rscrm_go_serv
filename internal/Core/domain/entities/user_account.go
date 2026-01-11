package entities

import "time"

type UserAccount struct {
	Id         int64
	Email      string
	Password   string
	Created_at time.Time
	School_id  string
}
