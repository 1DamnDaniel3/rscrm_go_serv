package accountrolespolicies

type AccountRolesPolicies struct {
	CRUD         IAccountRolesCrudPolicy
	CreatePolicy IAccountRolesCreatePolicy
	DeletePolicy IAccountRolesDeletePolicy
}

func NewAccountRolesPolicies(
	crud IAccountRolesCrudPolicy,
	CreatePolicy IAccountRolesCreatePolicy,
	DeletePolicy IAccountRolesDeletePolicy,
) *AccountRolesPolicies {
	return &AccountRolesPolicies{
		CRUD:         crud,
		CreatePolicy: CreatePolicy,
		DeletePolicy: DeletePolicy,
	}
}
