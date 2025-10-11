package dto

type StudentGroupCreateUpdateDTO struct {
	ID         int64  `json:"id"`
	Student_id int64  `json:"student_id"`
	Group_id   int64  `json:"group_id"`
	School_id  string `json:"school_id"`
}

type StudentGroupResponseDTO struct {
	ID         int64  `json:"id"`
	Student_id int64  `json:"student_id"`
	Group_id   int64  `json:"group_id"`
	School_id  string `json:"school_id"`
}
