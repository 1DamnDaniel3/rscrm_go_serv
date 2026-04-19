package policytypes

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"

type User struct {
	ID        int64
	Email     string
	Roles     []valuetypes.Role
	School_id string
}
