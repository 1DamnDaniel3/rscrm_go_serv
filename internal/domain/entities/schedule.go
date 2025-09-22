package entities

import "time"

type Schedule struct {
	ID               int64
	Group_id         int64
	Direction_id     int64
	Teacher_id       int64
	Weekday          int
	Start_time       time.Time
	Duration_minutes int
	School_id        string
	Active_from      time.Time
	Active_to        time.Time
}
