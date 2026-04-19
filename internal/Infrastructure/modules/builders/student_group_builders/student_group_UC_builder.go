package studentgroupbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type StudentGroupUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.StudentGroup]
}

func NewStudentGroupUseCasesBuilder(
	studentGroupModule *modules.StudentGroupsModule,
) *StudentGroupUseCases {

	return &StudentGroupUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			studentGroupModule.StudentGroupsRepo,
			studentGroupModule.StudentGroupPolicies.CRUD,
		),
	}
}
