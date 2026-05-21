package modules

import (
	dancestylespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/dance_styles_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type DanceStyleModule struct {
	DanceStyleRepo     entitiesrepos.DanceStylesRepo
	DanceStylePolicies *dancestylespolicies.DanceStylesPolicies
}

func NewDanceStyleModule(
	db *gorm.DB,
) *DanceStyleModule {
	return &DanceStyleModule{
		DanceStyleRepo: gormentityrepos.NewGormDanceStylesRepo(db),

		DanceStylePolicies: dancestylespolicies.NewDanceStylesPolicies(
			dancestylespolicies.NewDanceStylesCrudPolicy(),
		),
	}
}
