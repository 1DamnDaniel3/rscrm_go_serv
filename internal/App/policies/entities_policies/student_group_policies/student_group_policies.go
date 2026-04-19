package studentgrouppolicies

type StudentGroupPolicies struct {
	CRUD IStudentGroupCrudPolicy
}

func NewStudentGroupPolicies(crud IStudentGroupCrudPolicy) *StudentGroupPolicies {
	return &StudentGroupPolicies{
		CRUD: crud,
	}
}
