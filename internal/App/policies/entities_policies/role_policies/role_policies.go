package rolepolicies

type RolesPolicies struct {
	CRUD IRolesCrudPolicy
}

func NewRolesPolicies(crud IRolesCrudPolicy) *RolesPolicies {
	return &RolesPolicies{
		CRUD: crud,
	}
}
