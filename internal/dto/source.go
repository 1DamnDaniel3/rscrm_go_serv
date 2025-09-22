package dto

type SourceCreateUpdateDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type SourceResponseDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
