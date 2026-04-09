package dto

import "time"

type UserAccountResponseDTO struct {
	Id         int64     `json:"id"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
	School_id  string    `json:"school_id"`
}

type UserAccountCreateDTO struct {
	Id         int64     `json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
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
	Created_at *time.Time `json:"created_at,omitempty"`
	School_id  *string    `json:"school_id,omitempty"`
}
