package dto

type TransactionParticipantCreateDTO struct {
	Id             int64  `json:"id"`
	Transaction_id int64  `json:"transaction_id"`
	Role           string `json:"role"` //CHECK (role IN ('payer', 'beneficiary')),
	Entity_type    string `json:"entity_type"`
	Entity_id      int64  `json:"entity_id"`
}

type TransactionParticipantResponseDTO struct {
	Id             int64  `json:"id"`
	Transaction_id int64  `json:"transaction_id"`
	Role           string `json:"role"` //CHECK (role IN ('payer', 'beneficiary')),
	Entity_type    string `json:"entity_type"`
	Entity_id      int64  `json:"entity_id"`
}
