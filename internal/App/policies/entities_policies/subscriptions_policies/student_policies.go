package subscriptionpolicies

type SubscriptionPolicies struct {
	CRUD ISubscriptionCrudPolicy
}

func NewSubscriptionPolicies(crud ISubscriptionCrudPolicy) *SubscriptionPolicies {
	return &SubscriptionPolicies{
		CRUD: crud,
	}
}
