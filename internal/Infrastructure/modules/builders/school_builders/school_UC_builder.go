package schoolbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type SchoolUseCases struct {
	Crud genericcruduc.ICRUDUseCase[entities.School]
}

func NewSchoolUseCasesBuilder(
	schoolModule *modules.SchoolModule,
) *SchoolUseCases {

	return &SchoolUseCases{
		Crud: genericcruduc.NewCRUDUseCase(
			schoolModule.SchoolRepo,
			schoolModule.SchoolPolicies.CRUD,
		),
	}
}
