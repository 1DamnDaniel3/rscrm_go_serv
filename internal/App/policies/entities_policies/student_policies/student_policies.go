package studentpolicies

type StudentPolicies struct {
	CRUD IStudentCrudPolicy
}

func NewStudentPolicies(crud IStudentCrudPolicy) *StudentPolicies {
	return &StudentPolicies{
		CRUD: crud,
	}
}
