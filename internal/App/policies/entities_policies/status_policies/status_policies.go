package statuspolicies

type StatusPolicies struct {
	CRUD IStatusCrudPolicy
}

func NewStatusPolicies(crud IStatusCrudPolicy) *StatusPolicies {
	return &StatusPolicies{
		CRUD: crud,
	}
}
