package studentclientpolicies

type StudentClientPolicies struct {
	CRUD IStudentClientCrudPolicy
}

func NewStudentClientPolicies(crud IStudentClientCrudPolicy) *StudentClientPolicies {
	return &StudentClientPolicies{
		CRUD: crud,
	}
}
