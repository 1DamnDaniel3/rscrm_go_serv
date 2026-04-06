package clientgroupspolicies

type ClientGroupsPolicies struct {
	CRUD IClientGroupsCrudPolicy
}

func NewClientGroupsPolicies(crud IClientGroupsCrudPolicy) *ClientGroupsPolicies {
	return &ClientGroupsPolicies{
		CRUD: crud,
	}
}
