package schedulebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type ScheduleHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Schedule,
		dto.ScheduleCreateUpdateDTO,
		dto.ScheduleResponseDTO,
	]
}

func NewScheduleHandlerBuilder(
	uc *ScheduleUseCases,
) *ScheduleHandlerBuilder {

	return &ScheduleHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Schedule,
			dto.ScheduleCreateUpdateDTO,
			dto.ScheduleResponseDTO,
		](uc.CRUD),
	}
}
