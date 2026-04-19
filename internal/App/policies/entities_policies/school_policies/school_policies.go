package schoolpolicies

type SchoolPolicies struct {
	CRUD ISchoolCrudPolicy
}

func NewSchoolPolicies(
	crud ISchoolCrudPolicy,
) *SchoolPolicies {
	return &SchoolPolicies{
		CRUD: crud,
	}
}
