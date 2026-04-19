package studentbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs"
	studentclientUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs/student_clientUCs"
	studentgroupUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs/student_groupUCs"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type StudentUseCases struct {
	CRUD               genericcruduc.ICRUDUseCase[entities.Student]
	CreateAndGroup     studentgroupUCs.ICreateStudentUC
	StudentGroupCRUD   studentgroupUCs.IStudentGroupCRUDucs
	GroupedStudents    studentucs.IGroupedStudentsUC
	GetGroupsByStudent studentucs.IGetStudentGroupUC
	SearchStudents     studentucs.ISearchStudentsUC
	GetStudentClients  studentclientUCs.IGetStudentClientsUC
}

func NewStudentUseCasesBuilder(
	tx services.Transaction,
	studentModule *modules.StudentModule,
	studentGroupModule *modules.StudentGroupsModule,
) *StudentUseCases {

	return &StudentUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			studentModule.StudentRepo,
			studentModule.StudentPolicies.CRUD,
		),

		// ================= GROUP + CREATE =================
		CreateAndGroup: studentgroupUCs.NewCreateStudentUC(
			tx,
			studentModule.StudentRepo,
			studentGroupModule.StudentGroupsRepo,
			studentModule.StudentPolicies.CRUD,
			studentGroupModule.StudentGroupPolicies.CRUD,
		),

		StudentGroupCRUD: studentgroupUCs.NewStudentGroupCRUDucs(
			studentGroupModule.StudentGroupsRepo,
			studentGroupModule.StudentGroupPolicies.CRUD,
		),

		// ================= GROUPED =================
		GroupedStudents: studentucs.NewGroupedStudentsUC(
			studentModule.StudentRepo,
		),

		// ================= RELATIONS =================
		GetGroupsByStudent: studentucs.NewGetStudentGroupUC(
			studentModule.StudentQueryService,
		),

		SearchStudents: studentucs.NewSearchStudentsUC(
			studentModule.StudentQueryService,
		),

		GetStudentClients: studentclientUCs.NewGetStudentClientsUC(
			studentModule.StudentQueryService,
		),
	}
}
