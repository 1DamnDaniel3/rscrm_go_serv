package entities

import "time"

type UserProfile struct {
	ID        string
	AccountID int64
	Phone     string
	FullName  string
	BirthDate time.Time
}
