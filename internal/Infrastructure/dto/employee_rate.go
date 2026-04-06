package dto

import "time"

type EmployeeRateCreateUpdateDTO struct {
	ID          int64     `json:"id"`
	Employee_id int64     `json:"teacher_id"`
	Policy_id   int64     `json:"policy_id"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	School_id   string    `json:"school_id"`
}

type EmployeeRateResponseDTO struct {
	ID          int64     `json:"id"`
	Employee_id int64     `json:"teacher_id"`
	Policy_id   int64     `json:"policy_id"`
	Active_from time.Time `json:"active_from"`
	Active_to   time.Time `json:"active_to"`
	School_id   string    `json:"school_id"`
}
