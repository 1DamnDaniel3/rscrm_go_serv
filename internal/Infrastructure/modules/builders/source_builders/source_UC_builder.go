package sourcebuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type SourceUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.Source]
}

func NewSourceUseCasesBuilder(
	sourceModule *modules.SourceModule,
) *SourceUseCases {

	return &SourceUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			sourceModule.SourceRepo,
			sourceModule.SourcePolicies.CRUD,
		),
	}
}
