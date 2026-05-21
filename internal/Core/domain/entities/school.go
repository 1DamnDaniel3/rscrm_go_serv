package entities

import "time"

type School struct {
	ID         string
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
