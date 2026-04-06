package entities

import "time"

type EmployeeRate struct {
	ID          int64
	Employee_id int64
	Policy_id   int64
	Active_from time.Time
	Active_to   time.Time
	School_id   string
}
