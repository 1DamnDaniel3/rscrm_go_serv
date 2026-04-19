package studentclientbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	studentclienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentClientHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type StudentClientHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.StudentClient,
		dto.StudentClientCreateUpdateDTO,
		dto.StudentClientResponseDTO,
	]

	CreateRelationHandler *studentclienthandlers.CreateStudentClientRelHandler
}

func NewStudentClientHandlerBuilder(
	uc *StudentClientUseCases,
) *StudentClientHandlerBuilder {

	return &StudentClientHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.StudentClient,
			dto.StudentClientCreateUpdateDTO,
			dto.StudentClientResponseDTO,
		](uc.CRUD),

		// ================= RELATION =================
		CreateRelationHandler: studentclienthandlers.NewCreateStudentClientRelHandler(
			uc.CreateRelationUC,
		),
	}
}
