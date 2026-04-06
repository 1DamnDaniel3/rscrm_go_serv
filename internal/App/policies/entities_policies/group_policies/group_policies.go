package grouppolicies

type GroupPolicies struct {
	CRUD IGroupCrudPolicy
}

func NewGroupPolicies(crud IGroupCrudPolicy) *GroupPolicies {
	return &GroupPolicies{
		CRUD: crud,
	}
}
