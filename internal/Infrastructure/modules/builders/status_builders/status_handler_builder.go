package statusbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type StatusHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Status,
		dto.StatusCreateUpdateDTO,
		dto.StatusResponseDTO,
	]
}

func NewStatusHandlerBuilder(
	uc *StatusUseCases,
) *StatusHandlerBuilder {

	return &StatusHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Status,
			dto.StatusCreateUpdateDTO,
			dto.StatusResponseDTO,
		](uc.CRUD),
	}
}
