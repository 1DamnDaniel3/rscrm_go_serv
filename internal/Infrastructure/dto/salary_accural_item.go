package dto

type SalaryAccrualItemsCreateDTO struct {
	Id          int64  `json:"id"`
	Accrual_id  int64  `json:"accrual_id"`
	Source_type string `json:"source_type"` //CHECK (source_type IN ('fixed', 'per_student', 'per_lesson', 'percent', 'bonus')),
	Source_id   int64  `json:"source_id"`   //-- ссылка на lesson, policy, transaction, если нужно
	Amount      string `json:"amount"`      // NOT NULL, -- salary_accruals.amount = SUM(items)
	School_id   string `json:"school_id"`
}

type SalaryAccrualItemsResponseDTO struct {
	Id          int64  `json:"id"`
	Accrual_id  int64  `json:"accrual_id"`
	Source_type string `json:"source_type"` //CHECK (source_type IN ('fixed', 'per_student', 'per_lesson', 'percent', 'bonus')),
	Source_id   int64  `json:"source_id"`   //-- ссылка на lesson, policy, transaction, если нужно
	Amount      string `json:"amount"`      // NOT NULL, -- salary_accruals.amount = SUM(items)
	School_id   string `json:"school_id"`
}
