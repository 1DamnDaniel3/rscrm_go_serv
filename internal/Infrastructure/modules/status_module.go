package modules

import (
	statuspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/status_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type StatusModule struct {
	StatusRepo     entitiesrepos.StatusRepo
	StatusPolicies *statuspolicies.StatusPolicies
}

func NewStatusModule(
	db *gorm.DB,
) *StatusModule {
	return &StatusModule{
		StatusRepo: gormentityrepos.NewGormStatusRepository(db),

		StatusPolicies: statuspolicies.NewStatusPolicies(
			statuspolicies.NewStatusCrudPolicy(),
		),
	}
}
