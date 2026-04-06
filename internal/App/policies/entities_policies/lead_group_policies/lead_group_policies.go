package leadgrouppolicies

type LeadGroupPolicies struct {
	CRUD ILeadGroupCrudPolicy
}

func NewLeadGroupPolicies(crud ILeadGroupCrudPolicy) *LeadGroupPolicies {
	return &LeadGroupPolicies{
		CRUD: crud,
	}
}
