package bodtos

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"

type BoDTO_StudentClientsReponse struct {
	Relation_id int64 `json:"relation_id"`
	*dto.ClientResponseDTO
	Is_payer bool   `json:"is_payer"`
	Relation string `json:"relation"`
}
