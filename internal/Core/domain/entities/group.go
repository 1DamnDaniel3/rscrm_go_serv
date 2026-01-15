package entities

import "time"

type Group struct {
	ID          int64
	Name        string
	Entity_type string
	Created_at  time.Time
	School_id   string
}

func (l *Group) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
