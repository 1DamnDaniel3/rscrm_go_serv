package transactionsbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type TransactionsUseCases struct {
	GenericCRUD genericcruduc.ICRUDUseCase[entities.Transaction]
}

func NewSubscriptionUCBuilder(
	TransactionModule *modules.TransactionModule,
) *TransactionsUseCases {
	return &TransactionsUseCases{

		GenericCRUD: genericcruduc.NewCRUDUseCase(
			TransactionModule.TransactionRepo,
			TransactionModule.TransactionsPolicy.CRUD,
		),
	}
}
