package dto

type LeadGroupCreateUpdateDTO struct {
	ID        int64  `json:"id"`
	Lead_id   int64  `json:"lead_id"`
	Group_id  int64  `json:"group_id"`
	School_id string `json:"school_id"`
}

type LeadGroupResponseDTO struct {
	ID        int64  `json:"id"`
	Lead_id   int64  `json:"lead_id"`
	Group_id  int64  `json:"group_id"`
	School_id string `json:"school_id"`
}
