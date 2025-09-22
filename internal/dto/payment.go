package dto

import "time"

type PaymentCreateUpdateDTO struct {
	ID              int64     `json:"id"`
	Student_id      int64     `json:"student_id"`
	Subscription_id int64     `json:"subscription_id"`
	Amount          float64   `json:"amount"`
	Paid_at         time.Time `json:"paid_at"`
	Created_by      int64     `json:"created_by"`
	School_id       string    `json:"school_id"`
}

type PaymentResponseDTO struct {
	ID              int64     `json:"id"`
	Student_id      int64     `json:"student_id"`
	Subscription_id int64     `json:"subscription_id"`
	Amount          float64   `json:"amount"`
	Paid_at         time.Time `json:"paid_at"`
	Created_by      int64     `json:"created_by"`
	School_id       string    `json:"school_id"`
}
