package schedulebuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type ScheduleUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.Schedule]
}

func NewScheduleUseCasesBuilder(
	tx services.Transaction,
	scheduleModule *modules.ScheduleModule,
) *ScheduleUseCases {

	return &ScheduleUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			scheduleModule.ScheduleRepo,
			scheduleModule.SchedulePolicies.CRUD,
		),
	}
}
