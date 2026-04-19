package modules

import (
	schedulepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/schedule_policies"
	schoolpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/school_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type SchoolModule struct {
	SchoolRepo entitiesrepos.SchoolRepository

	SchoolPolicies *schoolpolicies.SchoolPolicies
}

func NewSchoolModule(
	db *gorm.DB,
) *SchoolModule {
	return &SchoolModule{
		SchoolRepo: gormentityrepos.NewGormSchoolRepo(db),

		SchoolPolicies: schoolpolicies.NewSchoolPolicies(
			schedulepolicies.NewScheduleCrudPolicy(),
		),
	}
}
