package useraccountpolicies

type UserAccountPolicies struct {
	CRUD       IUserAccountCrudPolicy
	ReadPolicy IUserAccountReadPolicy
}

func NewUserAccountPolicies(
	crud IUserAccountCrudPolicy,
	ReadPolicy IUserAccountReadPolicy,
) *UserAccountPolicies {
	return &UserAccountPolicies{
		CRUD:       crud,
		ReadPolicy: ReadPolicy,
	}
}
