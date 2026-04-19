package clientbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"

	clientgroupsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/clientUCs/client_groupUCs"
	clientstudentsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/clientUCs/client_studentUCs"

	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type ClientUseCases struct {
	CRUD                genericcruduc.ICRUDUseCase[entities.Client]
	GroupedClients      clientgroupsUCs.IGroupedClientsUC
	CreateGroupedClient clientgroupsUCs.ICreateGroupedClientUC
	GetClientGroups     clientgroupsUCs.IGetClientGroupUC
	GetClientStudents   clientstudentsUCs.IGetClientStudentsUC
}

func NewClientUseCases(
	tx services.Transaction,
	clientModule *modules.ClientModule,
	clientGroupsModule *modules.ClientGroupsModule,

) *ClientUseCases {

	return &ClientUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			clientModule.ClientRepo,
			clientModule.ClientPolicies.CRUD,
		),

		// ================= GROUPED =================
		GroupedClients: clientgroupsUCs.NewGroupedClientsUC(
			clientModule.ClientRepo,
		),

		CreateGroupedClient: clientgroupsUCs.NewCreateGroupedClientUC(
			tx,
			clientGroupsModule.ClientGroupRepo,
			clientModule.ClientRepo,
			clientGroupsModule.ClientGroupsPolicies.CRUD,
			clientModule.ClientPolicies.CRUD,
		),

		// ================= RELATIONS =================
		GetClientGroups: clientgroupsUCs.NewGetClientGroupUC(
			clientModule.ClientQueryService,
		),

		GetClientStudents: clientstudentsUCs.NewGetClientStudentsUC(
			clientModule.ClientQueryService,
		),
	}
}
