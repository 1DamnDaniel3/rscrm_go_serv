package valuetypes

type Role string

const (
	Admin        Role = "admin"
	Manager      Role = "manager"
	Teacher      Role = "teacher"
	Receptionist Role = "receptionist"
	Accountant   Role = "accountant"
	Owner        Role = "owner"
)
