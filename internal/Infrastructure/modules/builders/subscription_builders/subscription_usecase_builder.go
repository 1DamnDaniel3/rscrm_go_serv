package subscriptionbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type SubscriptionUseCases struct {
	GenericCRUD genericcruduc.ICRUDUseCase[entities.Subscription]
}

func NewSubscriptionUCBuilder(
	SubscriptionModule *modules.SubscriptionModule,
) *SubscriptionUseCases {
	return &SubscriptionUseCases{

		GenericCRUD: genericcruduc.NewCRUDUseCase(
			SubscriptionModule.SubscriptionRepo,
			SubscriptionModule.SubscriptionPolicies.CRUD,
		),
	}
}
