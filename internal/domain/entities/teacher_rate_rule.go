package entities

type TeacherRateRule struct {
	ID                int64
	Policy_id         int64
	Rule_type         string
	Threshold         int
	Base_amount       float64
	Per_student       float64
	Percent_of_income float64
	School_id         string
}
