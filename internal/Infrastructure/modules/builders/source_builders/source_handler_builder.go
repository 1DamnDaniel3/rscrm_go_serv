package sourcebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type SourceHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Source,
		dto.SourceCreateUpdateDTO,
		dto.SourceResponseDTO,
	]
}

func NewSourceHandlerBuilder(
	uc *SourceUseCases,
) *SourceHandlerBuilder {

	return &SourceHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Source,
			dto.SourceCreateUpdateDTO,
			dto.SourceResponseDTO,
		](uc.CRUD),
	}
}
