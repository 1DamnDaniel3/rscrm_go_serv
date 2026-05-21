package lessonsubscriptionsbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type LessonSubscriptionsUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.LessonSubscription]
}

func NewLessonSubscriptionsUseCases(
	tx services.Transaction,
	lessonSubscriptionModule *modules.LessonSubscriptionsModule,
) *LessonSubscriptionsUseCases {

	return &LessonSubscriptionsUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			lessonSubscriptionModule.LessonSubscriptionRepo,
			lessonSubscriptionModule.LessonSubscriptionPolicies.CRUD,
		),
	}
}
