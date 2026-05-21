package salaryaccuralspolicies

type SalaryAccuralsPolicies struct {
	CRUD ISalaryAccuralsCrudPolicy
}

func NewSalaryAccuralsPolicies(crud ISalaryAccuralsCrudPolicy) *SalaryAccuralsPolicies {
	return &SalaryAccuralsPolicies{
		CRUD: crud,
	}
}
