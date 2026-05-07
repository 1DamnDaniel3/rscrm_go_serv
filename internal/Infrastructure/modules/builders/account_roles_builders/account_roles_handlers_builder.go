package accountrolesbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	accountroleshandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user_handlers/account_roles_handlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type AccountRolesHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.AccountRoles,
		dto.AccountRolesCreateUpdateDTO,
		dto.AccountRolesResponseDTO,
	]
	NotGenericCrudHandler *accountroleshandlers.AccountRolesCrudHandler
}

func NewAccountRolesHandlerBuilder(useCases *AccountRolesUseCases) *AccountRolesHandlerBuilder {
	return &AccountRolesHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.AccountRoles,
			dto.AccountRolesCreateUpdateDTO,
			dto.AccountRolesResponseDTO,
		](useCases.CRUD),

		NotGenericCrudHandler: accountroleshandlers.NewAccountRolesCrudHandler(
			useCases.NotGenericCrud,
		),
	}
}
