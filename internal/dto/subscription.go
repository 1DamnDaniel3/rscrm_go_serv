package dto

import "time"

type SubscriptionCreateUpdateDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Visit_limit int       `json:"visit_limit"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	Is_archived bool      `json:"is_archived"`
	School_id   string    `json:"school_id"`
}

type SubscriptionResponseDTO struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price       float64   `json:"price"`
	Visit_limit int       `json:"visit_limit"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	Is_archived bool      `json:"is_archived"`
	School_id   string    `json:"school_id"`
}
