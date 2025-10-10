package entities

import "time"

type UserProfile struct {
	ID         int64
	Account_id int64
	Phone      string
	Full_name  string
	Birthdate  time.Time
}
