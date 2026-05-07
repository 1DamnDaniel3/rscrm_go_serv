package bodtos

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"

type BoDTO_User struct {
	User  dto.UserAccountResponseDTO `json:"account"`
	Roles []string                   `json:"roles"`
}
