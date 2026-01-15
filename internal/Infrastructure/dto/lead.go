package dto

import "time"

type LeadCreateUpdateDTO struct {
	ID                     int64      `json:"id"`
	Name                   string     `json:"name"`
	Phone                  string     `json:"phone"`
	Source_id              *int64     `json:"source_id"`
	Status_id              *int64     `json:"status_id"`
	Trial_date             *time.Time `json:"trial_date,omitempty"`
	Qualification          string     `json:"qualification"`
	Created_by             *int64     `json:"created_by"`
	Created_at             time.Time  `json:"created_at,omitempty"`
	Converted_to_client_at *time.Time `json:"converted_to_client_at,omitempty"`
	School_id              string     `json:"school_id"`
}

type LeadResponseDTO struct {
	ID                     int64      `json:"id"`
	Name                   string     `json:"name"`
	Phone                  string     `json:"phone"`
	Source_id              *int64     `json:"source_id"`
	Status_id              *int64     `json:"status_id"`
	Trial_date             *time.Time `json:"trial_date,omitempty"`
	Qualification          string     `json:"qualification"`
	Created_by             *int64     `json:"created_by"`
	Created_at             time.Time  `json:"created_at,omitempty"`
	Converted_to_client_at *time.Time `json:"converted_to_client_at,omitempty"`
	School_id              string     `json:"school_id"`
}
