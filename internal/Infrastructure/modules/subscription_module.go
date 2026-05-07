package modules

import (
	subscriptionpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/subscriptions_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type SubscriptionModule struct {
	SubscriptionRepo     entitiesrepos.SubscriptionRepo
	SubscriptionPolicies *subscriptionpolicies.SubscriptionPolicies
}

func NewSubscriptionModule(db *gorm.DB) *SubscriptionModule {
	return &SubscriptionModule{
		SubscriptionRepo: gormentityrepos.NewGormSubscriptionRepo(db),
		SubscriptionPolicies: subscriptionpolicies.NewSubscriptionPolicies(
			subscriptionpolicies.NewSubscriptionCrudPolicy(),
		),
	}
}
