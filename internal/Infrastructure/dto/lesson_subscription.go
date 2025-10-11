package dto

import "time"

type LessonSubscriptionCreateUpdateDTO struct {
	ID              int64     `json:"id"`
	Lesson_id       int64     `json:"lesson_id"`
	Student_id      int64     `json:"student_id"`
	Subscription_id int64     `json:"subscription_id"`
	Used_at         time.Time `json:"used_at"`
	School_id       string    `json:"school_id"`
}

type LessonSubscriptionResponseDTO struct {
	ID              int64     `json:"id"`
	Lesson_id       int64     `json:"lesson_id"`
	Student_id      int64     `json:"student_id"`
	Subscription_id int64     `json:"subscription_id"`
	Used_at         time.Time `json:"used_at"`
	School_id       string    `json:"school_id"`
}
