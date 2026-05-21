package useraccountbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/transactions"
	userhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user_handlers"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type UserAccountHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.UserAccount,
		dto.UserAccountCreateDTO,
		dto.UserAccountResponseDTO,
	]

	LoginHandler     *userhandlers.AuthHandler
	RegisterHandler  *transactions.RegisterHandler
	AuthCheckHandler *userhandlers.AuthCheckHandler
	GetMeHandler     *userhandlers.GetMeHandler

	GetAllAccountsWithRolesHandler *userhandlers.GetAllAccountsWithRolesHandler

	CreateEmployeeHandler *userhandlers.CreateEmployeeAccountHandler
}

func NewUserAccountHandlerBuilder(
	uc *UserUseCases,
	app *infrastructure.AppContainer,
) *UserAccountHandlerBuilder {

	return &UserAccountHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.UserAccount,
			dto.UserAccountCreateDTO,
			dto.UserAccountResponseDTO,
		](uc.CRUD),

		// ================= AUTH =================
		LoginHandler: userhandlers.NewLoginHandler(
			uc.Login,
		),

		RegisterHandler: transactions.NewRegisterHandler(
			uc.Register,
		),

		AuthCheckHandler: userhandlers.NewAuthCheckHandler(
			app.JWT,
			uc.CRUD,
		),

		// ================= GETME =================
		GetMeHandler: userhandlers.NewGetMeHandler(
			uc.GetMe,
		),

		// ================= GET ALL ACCS WITH ROLES =================
		GetAllAccountsWithRolesHandler: userhandlers.NewGetAllAccountsWithRolesHandler(
			uc.GetAllAccountsWithRolesUC,
			uc.GetAllAccountsWithRolesFilteredUC,
		),

		// ================= EMPLOYEE =================
		CreateEmployeeHandler: userhandlers.NewCreateEmployeeAccountHandler(
			uc.CreateEmployeeAcc,
		),
	}
}
