package modules

import (
	userprofilepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_profile_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type ProfileModule struct {
	ProfileRepo     entitiesrepos.UserProfileRepo
	ProfilePolicies *userprofilepolicies.UserProfilePolicies
}

func NewProfileModule(
	db *gorm.DB,
) *ProfileModule {
	return &ProfileModule{
		ProfileRepo: gormentityrepos.NewGormUserProfileRepo(db),

		ProfilePolicies: userprofilepolicies.NewUserProfilePolicy(
			userprofilepolicies.NewUserProfieCrudPolicy(),
			userprofilepolicies.NewReadProfilePolicy(),
		),
	}
}
