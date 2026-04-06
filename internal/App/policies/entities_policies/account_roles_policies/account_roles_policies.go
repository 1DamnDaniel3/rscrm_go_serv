package accountrolespolicies

type AccountRolesPolicies struct {
	CRUD IAccountRolesCrudPolicy
}

func NewAccountRolesPolicies(crud IAccountRolesCrudPolicy) *AccountRolesPolicies {
	return &AccountRolesPolicies{
		CRUD: crud,
	}
}
