package entities

import "time"

type UserProfile struct {
	ID         string
	Account_id int64
	Phone      string
	Full_name  string
	Birthdate  time.Time
}
