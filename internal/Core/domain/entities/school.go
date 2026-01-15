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

func (l *School) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
