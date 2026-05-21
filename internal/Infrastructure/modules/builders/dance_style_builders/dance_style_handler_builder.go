package dancestylebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type DanceStyleHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.DanceStyle,
		dto.DanceStyleCreateUpdateDTO,
		dto.DanceStyleResponseDTO,
	]
}

func NewClientHandlerBuilder(useCases *DanceStyleUseCases, modules *modules.DanceStyleModule) *DanceStyleHandlerBuilder {
	return &DanceStyleHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.DanceStyle,
			dto.DanceStyleCreateUpdateDTO,
			dto.DanceStyleResponseDTO,
		](useCases.CRUD),
	}
}
