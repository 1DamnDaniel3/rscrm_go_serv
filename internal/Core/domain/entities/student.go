package entities

import "time"

type Student struct {
	Id          int64
	Name        string
	Birthdate   *time.Time
	Skill_level string
	Contact     string
	Created_at  time.Time
	School_id   string
}

func (l *Student) BeforeCreate() error {
	if l.Created_at.IsZero() {
		l.Created_at = time.Now()
	}
	return nil
}
