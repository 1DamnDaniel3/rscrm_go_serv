package entities

import "time"

type Client struct {
	ID        int64
	Name      string
	Phone     string
	Birthdate time.Time
	Contact   string
	CreatedAt time.Time
	SchoolID  string
}
