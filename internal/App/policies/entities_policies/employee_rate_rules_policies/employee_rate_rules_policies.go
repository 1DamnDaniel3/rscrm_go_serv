package employeeraterulespolicies

type EmployeeRateRulesPolicies struct {
	CRUD IEmployeeRateRulesCrudPolicy
}

func NewEmployeeRateRulesPolicies(crud IEmployeeRateRulesCrudPolicy) *EmployeeRateRulesPolicies {
	return &EmployeeRateRulesPolicies{
		CRUD: crud,
	}
}
