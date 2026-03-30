package businessobjects

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

// =========== ClietnStudents
type GetClientStudentsBO struct {
	Relation_id int64
	entities.Student
	Relation string
}
