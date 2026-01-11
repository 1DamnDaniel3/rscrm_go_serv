package entities

import "time"

type UserProfile struct {
	Id         int64
	Account_id int64
	Phone      string
	Full_name  string
	Birthdate  time.Time
	School_id  string
}
