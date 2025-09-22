package entities

import "time"

type Student struct {
	ID          int64
	Name        string
	Birthdate   time.Time
	Skill_level string
	Contact     string
	Created_at  time.Time
	School_id   string
}
