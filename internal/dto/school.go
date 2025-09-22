package dto

import "time"

type SchoolCreateUpdateDTO struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}

type SchoolResponseDTO struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	City       string    `json:"city"`
	Phone      string    `json:"phone"`
	Email      string    `json:"email"`
	Created_at time.Time `json:"created_at"`
}
