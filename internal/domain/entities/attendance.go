package entities

import "time"

type Attendance struct {
	ID         int64
	Student_id int64
	Lesson_id  int64
	Status     string
	Marked_by  int64
	Marked_at  time.Time
	School_id  string
}
