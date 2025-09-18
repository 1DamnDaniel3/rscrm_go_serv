package entities

import "time"

type Student struct {
	ID         int64
	Name       string
	Birthdate  time.Time
	SkillLevel string
	Contact    string
	CreatedAt  time.Time
	SchoolID   string
}
