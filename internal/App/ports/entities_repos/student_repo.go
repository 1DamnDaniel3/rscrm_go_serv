package entitiesrepos

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type StudentsRepo interface {
	generic.Repository[entities.Student]
	GetGroupedStudents(ctx context.Context, group_id int64, entities *[]entities.Student) error
}

type StudentQueryService interface {
	GetStudentClients(ctx context.Context, student_id int64, clientsSlice *[]businessobjects.GetStudentClientsBO) error
}
