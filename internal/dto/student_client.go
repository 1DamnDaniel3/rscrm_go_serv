package dto

type StudentClientCreateUpdateDTO struct {
	ID         int64  `json:"id"`
	Student_id int64  `json:"student_id"`
	Client_id  int64  `json:"client_id"`
	Is_payer   bool   `json:"is_payer"`
	Relation   string `json:"relation"`
	School_id  string `json:"school_id"`
}

type StudentClientResponseDTO struct {
	ID         int64  `json:"id"`
	Student_id int64  `json:"student_id"`
	Client_id  int64  `json:"client_id"`
	Is_payer   bool   `json:"is_payer"`
	Relation   string `json:"relation"`
	School_id  string `json:"school_id"`
}
