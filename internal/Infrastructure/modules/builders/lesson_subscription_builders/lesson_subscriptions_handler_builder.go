package lessonsubscriptionsbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type LessonSubscriptionsHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.LessonSubscription,
		dto.LessonSubscriptionCreateUpdateDTO,
		dto.LessonSubscriptionResponseDTO,
	]
}

func NewLessonSubscriptionsHandlerBuilder(
	uc *LessonSubscriptionsUseCases,
) *LessonSubscriptionsHandlerBuilder {

	return &LessonSubscriptionsHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.LessonSubscription,
			dto.LessonSubscriptionCreateUpdateDTO,
			dto.LessonSubscriptionResponseDTO,
		](uc.CRUD),
	}
}
