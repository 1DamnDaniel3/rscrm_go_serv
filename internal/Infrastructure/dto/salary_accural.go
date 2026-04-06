package dto

import "time"

type SalaryAccuracyCreateDTO struct {
	Id              int64     `json:"id"`
	Employee_id     int64     `json:"employee_id"`
	Period_start    time.Time `json:"period_start"`
	Period_end      time.Time `json:"period_end"`
	Amount          string    `json:"amount"`
	Status          string    `json:"status"` //CHECK (status IN ('calculated', 'approved', 'paid')),
	Calculated_at   time.Time `json:"calculated_at"`
	TimeApproved_at time.Time `json:"time_approved_at"`
	School_id       string    `json:"school_id"`
}

type SalaryAccuracyResponseDTO struct {
	Id              int64     `json:"id"`
	Employee_id     int64     `json:"employee_id"`
	Period_start    time.Time `json:"period_start"`
	Period_end      time.Time `json:"period_end"`
	Amount          string    `json:"amount"`
	Status          string    `json:"status"` //CHECK (status IN ('calculated', 'approved', 'paid')),
	Calculated_at   time.Time `json:"calculated_at"`
	TimeApproved_at time.Time `json:"time_approved_at"`
	School_id       string    `json:"school_id"`
}
