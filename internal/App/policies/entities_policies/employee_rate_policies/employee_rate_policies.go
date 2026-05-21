package employeeratepolicies

type EmployeeRatePolicies struct {
	CRUD IEmployeeRateCrudPolicy
}

func NewEmployeeRatePolicyPolicies(crud IEmployeeRateCrudPolicy) *EmployeeRatePolicies {
	return &EmployeeRatePolicies{
		CRUD: crud,
	}
}
