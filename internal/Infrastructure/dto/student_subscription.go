package dto

import "time"

type StudentSubscriptionCreateUpdateDTO struct {
	ID               int64     `json:"id"`
	Student_id       int64     `json:"student_id"`
	Subscription_id  int64     `json:"subscription_id"`
	Issued_at        time.Time `json:"issued_at"`
	Expires_at       time.Time `json:"expires_at"`
	Remaining_visits int       `json:"remaining_visits"`
	Is_active        bool      `json:"is_active"`
	School_id        string    `json:"school_id"`
}

type StudentSubscriptionResponseDTO struct {
	ID               int64     `json:"id"`
	Student_id       int64     `json:"student_id"`
	Subscription_id  int64     `json:"subscription_id"`
	Issued_at        time.Time `json:"issued_at"`
	Expires_at       time.Time `json:"expires_at"`
	Remaining_visits int       `json:"remaining_visits"`
	Is_active        bool      `json:"is_active"`
	School_id        string    `json:"school_id"`
}
