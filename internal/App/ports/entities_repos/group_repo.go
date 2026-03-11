package entitiesrepos

import (
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupRepo interface {
	genericrepo.Repository[entities.Group]
}
