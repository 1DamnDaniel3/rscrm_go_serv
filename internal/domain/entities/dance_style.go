package entities

import "time"

type DanceStyle struct {
	ID         int64
	Name       string
	ActiveFrom time.Time
	ActiveTo   time.Time
	IsArchived bool
	SchoolID   string
}
