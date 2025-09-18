package entities

type TeacherRateRule struct {
	ID              int64
	PolicyID        int64
	RuleType        string
	Threshold       int
	BaseAmount      float64
	PerStudent      float64
	PercentOfIncome float64
	SchoolID        string
}
