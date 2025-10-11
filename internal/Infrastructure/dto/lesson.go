package dto

import "time"

type LessonCreateUpdateDTO struct {
	ID               int64     `json:"id"`
	Group_id         int64     `json:"group_id"`
	Direction_id     int64     `json:"direction_id"`
	Teacher_id       int64     `json:"teacher_id"`
	Lesson_date      time.Time `json:"lesson_date"`
	Start_time       time.Time `json:"start_time"`
	Duration_minutes int       `json:"duration_minutes"`
	Is_canceled      bool      `json:"is_canceled"`
	School_id        string    `json:"school_id"`
}

type LessonResponseDTO struct {
	ID               int64     `json:"id"`
	Group_id         int64     `json:"group_id"`
	Direction_id     int64     `json:"direction_id"`
	Teacher_id       int64     `json:"teacher_id"`
	Lesson_date      time.Time `json:"lesson_date"`
	Start_time       time.Time `json:"start_time"`
	Duration_minutes int       `json:"duration_minutes"`
	Is_canceled      bool      `json:"is_canceled"`
	School_id        string    `json:"school_id"`
}
