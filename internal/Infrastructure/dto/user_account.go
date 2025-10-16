package dto

import "time"

type UserAccountResponseDTO struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	School_id  string    `json:"school_id"`
}

type UserAccountCreateDTO struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	Created_at time.Time `json:"created_at"`
	School_id  string    `json:"school_id"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserAccountUpdateDTO struct {
	Email      *string    `json:"email,omitempty"`
	Password   *string    `json:"password,omitempty"`
	Role       *string    `json:"role,omitempty"`
	Created_at *time.Time `json:"created_at,omitempty"`
	School_id  *string    `json:"school_id,omitempty"`
}
