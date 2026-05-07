package dto

type AccountRolesCreateUpdateDTO struct {
	Id         int64  `json:"id"`
	Account_id int64  `json:"account_id"`
	Role_id    int64  `json:"role_id"`
	School_id  string `json:"school_id"`
}

type AccountRolesResponseDTO struct {
	Id         int64  `json:"id"`
	Account_id int64  `json:"account_id"`
	Role_id    int64  `json:"role_id"`
	School_id  string `json:"school_id"`
}
