package dto

import "time"

type FinancialReportCreateUpdateDTO struct {
	ID             int64     `json:"id"`
	Period_start   time.Time `json:"period_start"`
	Period_end     time.Time `json:"period_end"`
	Total_income   float64   `json:"total_income"`
	Total_expenses float64   `json:"total_expenses"`
	Created_at     time.Time `json:"created_at"`
	School_id      string    `json:"school_id"`
}

type FinancialReportResponseDTO struct {
	ID             int64     `json:"id"`
	Period_start   time.Time `json:"period_start"`
	Period_end     time.Time `json:"period_end"`
	Total_income   float64   `json:"total_income"`
	Total_expenses float64   `json:"total_expenses"`
	Created_at     time.Time `json:"created_at"`
	School_id      string    `json:"school_id"`
}
