package bodtos

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

// =========== ClietnStudents
type BoDTO_ClientStudentsReponse struct {
	Relation_id int64 `json:"relation_id"`
	StudentsAndGroups
	Relation string `json:"relation"`
}

type StudentsAndGroups struct {
	*dto.StudentResponseDTO
	Groups []dto.GroupResponseDTO `json:"groups"`
}
