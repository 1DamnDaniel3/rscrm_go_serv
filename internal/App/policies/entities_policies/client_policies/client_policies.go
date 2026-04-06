package employeeratepolicypolicies

type ClientPolicies struct {
	CRUD IClientCrudPolicy
}

func NewClientPolicies(crud IClientCrudPolicy) *ClientPolicies {
	return &ClientPolicies{
		CRUD: crud,
	}
}
