package studentsubscriptionbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type StudentSubscriptionsUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.StudentSubscription]
}

func NewStudentClientUseCasesBuilder(
	tx services.Transaction,
	studentSubscriptionModule *modules.StudentSubscriptionsModule,
) *StudentSubscriptionsUseCases {

	return &StudentSubscriptionsUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			studentSubscriptionModule.StudentSubscriptionsRepo,
			studentSubscriptionModule.StudentSubscriptionsPolicies.CRUD,
		),
	}
}
