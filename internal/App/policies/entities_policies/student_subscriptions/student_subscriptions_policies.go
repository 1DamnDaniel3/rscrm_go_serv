package studentsubscriptionpolicies

type StudentSubscriptionsPolicies struct {
	CRUD IStudentSubscriptionsCrudPolicy
}

func NewStudentPolicies(crud IStudentSubscriptionsCrudPolicy) *StudentSubscriptionsPolicies {
	return &StudentSubscriptionsPolicies{
		CRUD: crud,
	}
}
