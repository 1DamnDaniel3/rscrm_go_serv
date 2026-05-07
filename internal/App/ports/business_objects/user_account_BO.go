package businessobjects

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

type UserBO struct {
	UserAccount entities.UserAccount
	Roles       []string
}
