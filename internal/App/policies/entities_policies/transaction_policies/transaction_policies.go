package transactionpolicies

type TransactionPolicies struct {
	CRUD ITransactionCrudPolicy
}

func NewTransactionPolicies(crud ITransactionCrudPolicy) *TransactionPolicies {
	return &TransactionPolicies{
		CRUD: crud,
	}
}
