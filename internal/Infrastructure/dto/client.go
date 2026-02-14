package dto

import "time"

type ClientCreateUpdateDTO struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Birthdate  *time.Time `json:"birthdate"`
	Contact    string     `json:"contact"`
	Created_at time.Time  `json:"created_at"`
	School_id  string     `json:"school_id"`
}

type ClientResponseDTO struct {
	ID         int64      `json:"id"`
	Name       string     `json:"name"`
	Phone      string     `json:"phone"`
	Birthdate  *time.Time `json:"birthdate"`
	Contact    string     `json:"contact"`
	Created_at time.Time  `json:"created_at"`
	School_id  string     `json:"school_id"`
}
