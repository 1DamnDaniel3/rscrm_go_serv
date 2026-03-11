package businessobjects

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

// =========== StudentClietns
type GetStudentClientsBO struct {
	Relation_id int64
	entities.Client
	Is_payer bool
	Relation string
}
