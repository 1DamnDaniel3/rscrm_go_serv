package transactionsbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type TransactionHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Transaction,
		dto.TransactionCreateDTO,
		dto.TransactionResponseDTO,
	]
}

func NewSubscriptionHandlerBuilder(
	uc *TransactionsUseCases,
) *TransactionHandlerBuilder {
	return &TransactionHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Transaction,
			dto.TransactionCreateDTO,
			dto.TransactionResponseDTO,
		](uc.GenericCRUD),
	}
}
