package dto

type EmployeeRateRuleCreateUpdateDTO struct {
	ID                int64  `json:"id"`
	Policy_id         int64  `json:"policy_id"`
	Rule_type         string `json:"rule_type"`
	Threshold         string `json:"threshold"`
	Base_amount       string `json:"base_amount"`
	Per_student       string `json:"per_student"`
	Percent_of_income string `json:"percent_of_income"`
	School_id         string `json:"school_id"`
}

type EmployeeRateRuleResponseDTO struct {
	ID                int64  `json:"id"`
	Policy_id         int64  `json:"policy_id"`
	Rule_type         string `json:"rule_type"`
	Threshold         string `json:"threshold"`
	Base_amount       string `json:"base_amount"`
	Per_student       string `json:"per_student"`
	Percent_of_income string `json:"percent_of_income"`
	School_id         string `json:"school_id"`
}
