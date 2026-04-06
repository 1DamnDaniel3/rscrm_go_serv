package useraccountpolicies

type UserAccountPolicies struct {
	CRUD IUserAccountCrudPolicy
}

func NewUserAccountPolicies(crud IUserAccountCrudPolicy) *UserAccountPolicies {
	return &UserAccountPolicies{
		CRUD: crud,
	}
}
