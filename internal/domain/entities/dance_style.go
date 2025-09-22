package entities

import "time"

type DanceStyle struct {
	ID          int64
	Name        string
	Active_from time.Time
	Active_to   time.Time
	IsArchived  bool
	School_id   string
}
