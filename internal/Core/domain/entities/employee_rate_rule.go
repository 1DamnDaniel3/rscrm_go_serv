package entities

import "github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"

type EmployeeRateRule struct {
	ID                int64
	Policy_id         int64
	Rule_type         string
	Threshold         valuetypes.Money
	Base_amount       valuetypes.Money
	Per_student       valuetypes.Money
	Percent_of_income string
	School_id         string
}
