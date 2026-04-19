package modules

import (
	schedulepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/schedule_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type ScheduleModule struct {
	ScheduleRepo     entitiesrepos.ScheduleRepo
	SchedulePolicies *schedulepolicies.SchedulePolicies
}

func NewScheduleModule(
	db *gorm.DB,
) *ScheduleModule {
	return &ScheduleModule{
		ScheduleRepo: gormentityrepos.NewGormScheduleRepo(db),

		SchedulePolicies: schedulepolicies.NewSchedulePolicies(
			schedulepolicies.NewScheduleCrudPolicy(),
		),
	}
}
