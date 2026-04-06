package userprofilepolicies

type UserProfilePolicies struct {
	CRUD         IUserProfieCrudPolicy
	ReadPolicies IReadProfilePolicy
}

func NewUserProfilePolicy(
	crud IUserProfieCrudPolicy,
	readPolicies IReadProfilePolicy) *UserProfilePolicies {
	return &UserProfilePolicies{
		CRUD:         crud,
		ReadPolicies: readPolicies,
	}
}
