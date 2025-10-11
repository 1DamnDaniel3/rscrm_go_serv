package dto

import "time"

type UserAccountCreateUpdateDTO struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	School_id  string    `json:"school_id"`
}

type UserAccountResponseDTO struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	School_id  string    `json:"school_id"`
}
