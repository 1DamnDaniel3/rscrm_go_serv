package dto

import "time"

type DanceStyleCreateUpdateDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	IsArchived  bool      `json:"is_archived"`
	School_id   string    `json:"school_id"`
}

type DanceStyleResponseDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	IsArchived  bool      `json:"is_archived"`
	School_id   string    `json:"school_id"`
}
