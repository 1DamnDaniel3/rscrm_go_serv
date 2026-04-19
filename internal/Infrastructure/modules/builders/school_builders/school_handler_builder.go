package schoolbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type SchoolHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.School,
		dto.SchoolCreateUpdateDTO,
		dto.SchoolResponseDTO,
	]
}

func NewSchoolHandlerBuilder(
	uc *SchoolUseCases,
) *SchoolHandlerBuilder {

	return &SchoolHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.School,
			dto.SchoolCreateUpdateDTO,
			dto.SchoolResponseDTO,
		](uc.Crud),
	}
}
