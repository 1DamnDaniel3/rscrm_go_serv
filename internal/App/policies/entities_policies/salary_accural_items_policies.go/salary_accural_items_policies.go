package salaryaccuralitemspolicies

type SalaryAccuralItemsPolicies struct {
	CRUD ISalaryAccuralItemsCrudPolicy
}

func NewSalaryAccuralItemsPolicies(crud ISalaryAccuralItemsCrudPolicy) *SalaryAccuralItemsPolicies {
	return &SalaryAccuralItemsPolicies{
		CRUD: crud,
	}
}
