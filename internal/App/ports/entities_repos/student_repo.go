package entitiesrepos

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type StudentsRepo interface {
	generic.Repository[entities.Student]
	// GetGroupedLeads(school_id string, group_id int64, entities *[]entities.Lead) error
}
