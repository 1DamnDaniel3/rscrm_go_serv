package dto

type ClientGroupCreateUpdateDTO struct {
	ID        int64  `json:"id"`
	Client_id int64  `json:"client_id"`
	Group_id  int64  `json:"group_id"`
	School_id string `json:"school_id"`
}

type ClientGroupResponseDTO struct {
	ID        int64  `json:"id"`
	Client_id int64  `json:"client_id"`
	Group_id  int64  `json:"group_id"`
	School_id string `json:"school_id"`
}
