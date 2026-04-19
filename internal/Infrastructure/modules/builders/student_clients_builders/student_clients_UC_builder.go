package studentclientbuilders

import (
	studentclientUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs/student_clientUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type StudentClientUseCases struct {
	CRUD             genericcruduc.ICRUDUseCase[entities.StudentClient]
	CreateRelationUC studentclientUCs.ICreateRelationUC
}

func NewStudentClientUseCasesBuilder(
	tx services.Transaction,
	studentClientModule *modules.StudentClientModule,
	clientModule *modules.ClientModule,
) *StudentClientUseCases {

	return &StudentClientUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			studentClientModule.StudentClientsRepo,
			studentClientModule.StudentClientPolicies.CRUD,
		),

		CreateRelationUC: studentclientUCs.NewCreateRelationUC(
			tx,
			studentClientModule.StudentClientsRepo,
			clientModule.ClientRepo,
			studentClientModule.StudentClientPolicies.CRUD,
			clientModule.ClientPolicies.CRUD,
		),
	}
}
