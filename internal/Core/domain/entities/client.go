package entities

import "time"

type Client struct {
	ID         int64
	Name       string
	Phone      string
	Birthdate  *time.Time
	Contact    string
	Created_at time.Time
	School_id  string
}

func (l *Client) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
