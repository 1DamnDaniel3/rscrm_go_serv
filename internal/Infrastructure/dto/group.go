package dto

import "time"

type GroupCreateUpdateDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Entity_type string    `json:"entity_type"`
	Created_at  time.Time `json:"created_at"`
	School_id   string    `json:"school_id"`
}

type GroupResponseDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Entity_type string    `json:"entity_type"`
	Created_at  time.Time `json:"created_at"`
	School_id   string    `json:"school_id"`
}
