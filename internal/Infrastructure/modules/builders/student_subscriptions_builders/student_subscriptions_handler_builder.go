package studentsubscriptionbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	studentclienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentClientHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type StudentSubscriptionsHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.StudentSubscription,
		dto.StudentSubscriptionCreateUpdateDTO,
		dto.StudentSubscriptionResponseDTO,
	]

	CreateRelationHandler *studentclienthandlers.CreateStudentClientRelHandler
}

func NewStudentSubscriptionsHandlerBuilder(
	uc *StudentSubscriptionsUseCases,
) *StudentSubscriptionsHandlerBuilder {

	return &StudentSubscriptionsHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.StudentSubscription,
			dto.StudentSubscriptionCreateUpdateDTO,
			dto.StudentSubscriptionResponseDTO,
		](uc.CRUD),
	}
}
