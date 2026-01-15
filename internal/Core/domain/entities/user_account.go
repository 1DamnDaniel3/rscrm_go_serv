package entities

import "time"

type UserAccount struct {
	Id         int64
	Email      string
	Password   string
	Created_at time.Time
	School_id  string
}

func (l *UserAccount) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
