package entitiesrepos

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type StudentGroupsRepo interface {
	generic.Repository[entities.StudentGroup]
}
