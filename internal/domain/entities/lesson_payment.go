package entities

import "time"

type LessonPayment struct {
	ID         int64
	Lesson_id  int64
	Student_id int64
	Amount     float64
	Paid_at    time.Time
	School_id  string
}
