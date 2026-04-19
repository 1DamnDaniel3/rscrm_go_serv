package modules

import (
	sourcepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/source_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type SourceModule struct {
	SourceRepo     entitiesrepos.SourceRepo
	SourcePolicies *sourcepolicies.SourcePolicies
}

func NewSourceModule(
	db *gorm.DB,
) *SourceModule {
	return &SourceModule{
		SourceRepo: gormentityrepos.NewGormSourceRepository(db),

		SourcePolicies: sourcepolicies.NewSourcePolicies(
			sourcepolicies.NewSourceCrudPolicy(),
		),
	}
}
