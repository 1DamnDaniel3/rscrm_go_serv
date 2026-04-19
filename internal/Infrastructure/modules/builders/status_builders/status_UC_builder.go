package statusbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type StatusUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.Status]
}

func NewStatusUseCasesBuilder(
	statusModule *modules.StatusModule,
) *StatusUseCases {

	return &StatusUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			statusModule.StatusRepo,
			statusModule.StatusPolicies.CRUD,
		),
	}
}
