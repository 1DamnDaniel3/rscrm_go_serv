package entities

import "time"

type Group struct {
	ID         int64
	Name       string
	EntityType string
	CreatedAt  time.Time
	SchoolID   string
}
