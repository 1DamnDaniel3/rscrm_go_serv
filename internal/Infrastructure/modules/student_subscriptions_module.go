package modules

import (
	studentsubscriptionpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_subscriptions"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type StudentSubscriptionsModule struct {
	StudentSubscriptionsRepo     entitiesrepos.StudentSubscriptionsRepo
	StudentSubscriptionsPolicies *studentsubscriptionpolicies.StudentSubscriptionsPolicies
}

func NewStudentSubscriptionsModule(db *gorm.DB) *StudentSubscriptionsModule {
	return &StudentSubscriptionsModule{
		StudentSubscriptionsRepo: gormentityrepos.NewGormStudentSubscriptionsRepo(db),
		StudentSubscriptionsPolicies: studentsubscriptionpolicies.NewStudentPolicies(
			studentsubscriptionpolicies.NewStudentCrudPolicy(),
		),
	}
}
