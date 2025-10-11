package dto

type TeacherRateRuleCreateUpdateDTO struct {
	ID                int64   `json:"id"`
	Policy_id         int64   `json:"policy_id"`
	Rule_type         string  `json:"rule_type"`
	Threshold         int     `json:"threshold"`
	Base_amount       float64 `json:"base_amount"`
	Per_student       float64 `json:"per_student"`
	Percent_of_income float64 `json:"percent_of_income"`
	School_id         string  `json:"school_id"`
}

type TeacherRateRuleResponseDTO struct {
	ID                int64   `json:"id"`
	Policy_id         int64   `json:"policy_id"`
	Rule_type         string  `json:"rule_type"`
	Threshold         int     `json:"threshold"`
	Base_amount       float64 `json:"base_amount"`
	Per_student       float64 `json:"per_student"`
	Percent_of_income float64 `json:"percent_of_income"`
	School_id         string  `json:"school_id"`
}
