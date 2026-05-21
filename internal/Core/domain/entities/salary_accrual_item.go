package entities

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"

type SalaryAccrualItems struct {
	ID          int64
	Accrual_id  int64
	Source_type string           //CHECK (source_type IN ('fixed', 'per_student', 'per_lesson', 'percent', 'bonus')),
	Source_id   int64            //-- ссылка на lesson, policy, transaction, если нужно
	Amount      valuetypes.Money // NOT NULL, -- salary_accruals.amount = SUM(items)
	School_id   string
}
