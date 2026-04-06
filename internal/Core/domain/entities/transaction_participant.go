package entities

type TransactionParticipant struct {
	Id             int64
	Transaction_id int64
	Role           string //CHECK (role IN ('payer', 'beneficiary')),
	Entity_type    string
	Entity_id      int64
}
