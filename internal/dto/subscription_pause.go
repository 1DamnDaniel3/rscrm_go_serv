package dto

import "time"

type SubscriptionPauseCreateUpdateDTO struct {
	ID                     int64     `json:"id"`
	StudentSubscription_id int64     `json:"student_subscription_id"`
	Paused_from            time.Time `json:"paused_from"`
	Paused_to              time.Time `json:"paused_to"`
	School_id              string    `json:"school_id"`
}

type SubscriptionPauseResponseDTO struct {
	ID                     int64     `json:"id"`
	StudentSubscription_id int64     `json:"student_subscription_id"`
	Paused_from            time.Time `json:"paused_from"`
	Paused_to              time.Time `json:"paused_to"`
	School_id              string    `json:"school_id"`
}
