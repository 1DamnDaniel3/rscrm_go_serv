package entitiesrepos

import (
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type EmployeeRateRuleRepo interface {
	genericrepo.Repository[entities.EmployeeRateRule]
}
