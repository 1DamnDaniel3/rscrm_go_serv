package leadpolicies

type LeadPolicies struct {
	CRUD ILeadCrudPolicy
}

func NewLeadPolicies(crud ILeadCrudPolicy) *LeadPolicies {
	return &LeadPolicies{
		CRUD: crud,
	}
}
