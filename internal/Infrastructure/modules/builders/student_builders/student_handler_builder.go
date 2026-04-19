package studentbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	studenthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers"
	studentclientHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers/studentClientHandlers.go"
	studentgroupHandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentHandlers/studentGroupHandlers.go"
	studentgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_group_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type StudentHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Student,
		dto.StudentCreateUpdateDTO,
		dto.StudentResponseDTO,
	]

	// ===== queries =====
	GroupedStudentsHandler *studenthandlers.GetGroupedStudentsHandler
	GetGroupsByStudent     *studenthandlers.GetStudentGroupsHandler
	SearchStudentsHandler  *studenthandlers.SearchStudentHandler

	// ===== relations =====
	StudentGroupHandler   *studentgroupHandlers.StudentGroupHandler
	StudentClientsHandler *studentclientHandlers.StudentClientsHandler
}

func NewStudentHandlerBuilder(
	uc *StudentUseCases,
	studentGroupUC *studentgroupbuilders.StudentGroupUseCases,
) *StudentHandlerBuilder {

	return &StudentHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Student,
			dto.StudentCreateUpdateDTO,
			dto.StudentResponseDTO,
		](uc.CRUD),

		// ================= GROUPED =================
		GroupedStudentsHandler: studenthandlers.NewGetGroupedStudentsHandler(
			uc.GroupedStudents,
		),

		// ================= SEARCH =================
		SearchStudentsHandler: studenthandlers.NewSearchStudentHandler(
			uc.SearchStudents,
		),

		// ================= GROUPS =================
		GetGroupsByStudent: studenthandlers.NewGetStudentGroupsHandler(
			uc.GetGroupsByStudent,
		),

		// ================= RELATIONS =================
		StudentGroupHandler: studentgroupHandlers.NewStudentGroupHandler(
			uc.CreateAndGroup,
			uc.StudentGroupCRUD,
			studentGroupUC.CRUD,
		),

		StudentClientsHandler: studentclientHandlers.NewStudentClientsHandler(
			uc.GetStudentClients,
		),
	}
}
