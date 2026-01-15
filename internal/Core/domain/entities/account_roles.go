package entities

type AccountRoles struct {
	Id         int64
	Account_id int64
	Role_id    int64
	School_id  string
}

// role codes
// 'admin' id code = 1
// 'owner' id code = 2
// 'manager' id code = 3
// 'teacher' id code = 4
// 'accountant' id code = 5
// 'receptionist' id code = 6
