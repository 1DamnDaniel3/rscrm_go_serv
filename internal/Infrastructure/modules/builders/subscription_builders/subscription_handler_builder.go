package subscriptionbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type SubscriptionHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Subscription,
		dto.SubscriptionCreateUpdateDTO,
		dto.SubscriptionResponseDTO,
	]
}

func NewSubscriptionHandlerBuilder(
	uc *SubscriptionUseCases,
) *SubscriptionHandlerBuilder {
	return &SubscriptionHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Subscription,
			dto.SubscriptionCreateUpdateDTO,
			dto.SubscriptionResponseDTO,
		](uc.GenericCRUD),
	}
}
