package groupbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type GroupHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Group,
		dto.GroupCreateUpdateDTO,
		dto.GroupResponseDTO,
	]
}

func NewGroupHandlerBuilder(useCases *GroupUseCases) *GroupHandlerBuilder {
	return &GroupHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Group,
			dto.GroupCreateUpdateDTO,
			dto.GroupResponseDTO,
		](useCases.CRUD),
	}
}
