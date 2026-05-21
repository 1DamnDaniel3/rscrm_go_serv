package modules

import (
	transactionpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/transaction_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type TransactionModule struct {
	TransactionRepo    entitiesrepos.TransactionRepo
	TransactionsPolicy *transactionpolicies.TransactionPolicies
}

func NewTransactionModule(db *gorm.DB) *TransactionModule {
	return &TransactionModule{
		TransactionRepo: gormentityrepos.NewGormTransactionRepo(db),
		TransactionsPolicy: transactionpolicies.NewTransactionPolicies(
			transactionpolicies.NewTransactionCrudPolicy(),
		),
	}
}
