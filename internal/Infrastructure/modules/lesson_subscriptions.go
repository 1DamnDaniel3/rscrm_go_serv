package modules

import (
	lessonsubscriptionspolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lesson_subscriptions_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type LessonSubscriptionsModule struct {
	LessonSubscriptionRepo     entitiesrepos.LessonSubscriptionRepo
	LessonSubscriptionPolicies *lessonsubscriptionspolicies.LessonSubscriptionsPolicies
}

func NewLessonSubscriptionsModule(
	db *gorm.DB,
) *LessonSubscriptionsModule {
	return &LessonSubscriptionsModule{
		LessonSubscriptionRepo: gormentityrepos.NewGormLessonSubscriptionRepo(db),
		LessonSubscriptionPolicies: lessonsubscriptionspolicies.NewLessonSubscriptionsPolicies(
			lessonsubscriptionspolicies.NewLessonSubscriptionsCrudPolicy(),
		),
	}
}
