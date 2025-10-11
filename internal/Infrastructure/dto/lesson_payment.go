package dto

import "time"

type LessonPaymentCreateUpdateDTO struct {
	ID         int64     `json:"id"`
	Lesson_id  int64     `json:"lesson_id"`
	Student_id int64     `json:"student_id"`
	Amount     float64   `json:"amount"`
	Paid_at    time.Time `json:"paid_at"`
	School_id  string    `json:"school_id"`
}

type LessonPaymentResponseDTO struct {
	ID         int64     `json:"id"`
	Lesson_id  int64     `json:"lesson_id"`
	Student_id int64     `json:"student_id"`
	Amount     float64   `json:"amount"`
	Paid_at    time.Time `json:"paid_at"`
	School_id  string    `json:"school_id"`
}
