package dto

type StatusCreateUpdateDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	School_id string `json:"school_id"`
}

type StatusResponseDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	School_id string `json:"school_id"`
}
