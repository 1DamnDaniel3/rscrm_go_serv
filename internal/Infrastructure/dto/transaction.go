package dto

import (
	"time"
)

type TransactionCreateDTO struct {
	ID             int64     `json:"id"`
	Type           string    `json:"type"`     // CHECK (type IN ('income', 'expense')),
	Category       string    `json:"category"` // lesson or subscription_payment or rent
	Amount         string    `json:"amount"`
	Reference_type string    `json:"reference_type"` // lesson/subscription
	Reference_id   int64     `json:"reference_id"`
	Create_at      time.Time `json:"createTime"`
	Created_by     int64     `json:"created_by"`
	School_id      string    `json:"school_id"`
}

type TransactionResponseDTO struct {
	ID             int64     `json:"id"`
	Type           string    `json:"type"`     // CHECK (type IN ('income', 'expense')),
	Category       string    `json:"category"` // lesson or subscription_payment or rent
	Amount         string    `json:"amount"`
	Reference_type string    `json:"reference_type"` // lesson/subscription
	Reference_id   int64     `json:"reference_id"`
	Create_at      time.Time `json:"createTime"`
	Created_by     int64     `json:"created_by"`
	School_id      string    `json:"school_id"`
}
