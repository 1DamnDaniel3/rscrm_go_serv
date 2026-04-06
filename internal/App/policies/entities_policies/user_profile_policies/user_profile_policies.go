package userprofilepolicies

type UserProfilePolicies struct {
	CRUD IUserProfieCrudPolicy
}

func NewUserProfilePolicy(crud IUserProfieCrudPolicy) *UserProfilePolicies {
	return &UserProfilePolicies{
		CRUD: crud,
	}
}
