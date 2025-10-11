package dto

import "time"

type TeacherRatePolicyCreateUpdateDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	School_id   string    `json:"school_id"`
}

type TeacherRatePolicyResponseDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created_at  time.Time `json:"created_at"`
	School_id   string    `json:"school_id"`
}
