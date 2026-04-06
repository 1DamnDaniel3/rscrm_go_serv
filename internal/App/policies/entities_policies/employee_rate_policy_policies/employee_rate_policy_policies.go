package employeeratepolicypolicies

type EmployeeRatePolicyPolicies struct {
	CRUD IEmployeeRatePolicyCrudPolicy
}

func NewEmployeeRatePolicyPolicies(crud IEmployeeRatePolicyCrudPolicy) *EmployeeRatePolicyPolicies {
	return &EmployeeRatePolicyPolicies{
		CRUD: crud,
	}
}
