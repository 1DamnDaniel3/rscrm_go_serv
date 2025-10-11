package entities

import "time"

type LessonSubscription struct {
	ID              int64
	Lesson_id       int64
	Student_id      int64
	Subscription_id int64
	Used_at         time.Time
	School_id       string
}
