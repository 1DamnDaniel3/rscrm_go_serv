package dto

import "time"

type AttendanceCreateUpdateDTO struct {
	ID         int64     `json:"id"`
	Student_id int64     `json:"student_id"`
	Lesson_id  int64     `json:"lesson_id"`
	Status     string    `json:"status"`
	Marked_by  int64     `json:"marked_by"`
	Marked_at  time.Time `json:"marked_at"`
	School_id  string    `json:"school_id"`
}

type AttendanceResponseDTO struct {
	ID         int64     `json:"id"`
	Student_id int64     `json:"student_id"`
	Lesson_id  int64     `json:"lesson_id"`
	Status     string    `json:"status"`
	Marked_by  int64     `json:"marked_by"`
	Marked_at  time.Time `json:"marked_at"`
	School_id  string    `json:"school_id"`
}
