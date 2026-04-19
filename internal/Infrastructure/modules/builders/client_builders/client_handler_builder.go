package clientbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	clienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/clientHandlers"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type ClientHandlerBuilder struct {
	CRUDHandler           *generichandler.GenericHandler[entities.Client, dto.ClientCreateUpdateDTO, dto.ClientResponseDTO]
	GroupedClientsHandler *clienthandlers.GroupedClientsHandler
	CreateGroupedHandler  *clienthandlers.CreateGroupedClientHandler
	GetClientGroups       *clienthandlers.GetClientGroupsHandler
	GetClientStudents     *clienthandlers.GetClientStudentsHandler
	SearchHandler         *clienthandlers.ClientSearchHandler
}

func NewClientHandlerBuilder(useCases *ClientUseCases, modules *modules.ClientModule) *ClientHandlerBuilder {
	return &ClientHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Client,
			dto.ClientCreateUpdateDTO,
			dto.ClientResponseDTO,
		](useCases.CRUD),

		GroupedClientsHandler: clienthandlers.NewGroupedClientsHandler(useCases.GroupedClients),

		CreateGroupedHandler: clienthandlers.NewCreateGroupedClientHandler(useCases.CreateGroupedClient),

		GetClientGroups: clienthandlers.NewGetClientGroupsHandler(useCases.GetClientGroups),

		GetClientStudents: clienthandlers.NewGetClientStudentsHandler(useCases.GetClientStudents),

		SearchHandler: clienthandlers.NewClientSearchHandler(modules.ClientQueryService),
	}
}
