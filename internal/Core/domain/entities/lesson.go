package entities

import "time"

type Lesson struct {
	ID               int64
	Group_id         int64
	Direction_id     int64
	Teacher_id       int64
	Lesson_date      time.Time
	Start_time       time.Time
	Duration_minutes int
	Is_canceled      bool
	School_id        string
}
