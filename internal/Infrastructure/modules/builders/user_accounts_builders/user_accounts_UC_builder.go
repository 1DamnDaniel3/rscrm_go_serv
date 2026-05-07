package useraccountbuilders

import (
	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	userucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type UserUseCases struct {
	CRUD                      genericcruduc.ICRUDUseCase[entities.UserAccount]
	Login                     userucs.ILoginUC
	Register                  userucs.IRegisterUseCase
	GetMe                     userucs.IGetMeUC
	CreateEmployeeAcc         userucs.ICreateEmployeeAccountUC
	GetAllAccountsWithRolesUC userucs.IGetAllAccountsWithRolesUC
}

func NewUserUseCasesBuilder(
	tx services.Transaction,
	hasher ports.PasswordHasher,
	jwt ports.JWTservice,

	userModule *modules.AccountModule,
	profileModule *modules.ProfileModule,
	accountRolesModule *modules.AccountRolesModule,
	rolesModule *modules.RolesModule,
	schoolModule *modules.SchoolModule,
) *UserUseCases {

	return &UserUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			userModule.UserRepo,
			userModule.AccountPolicies.CRUD,
		),

		// ================= LOGIN =================
		Login: userucs.NewLoginUseCase(
			hasher,
			userModule.UserRepo,
			accountRolesModule.AccountRolesRepo,
			rolesModule.RolesRepo,
			jwt,
		),

		// ================= REGISTER =================
		Register: userucs.NewRegisterUseCase(
			tx,
			userModule.UserRepo,
			profileModule.ProfileRepo,
			schoolModule.SchoolRepo,
			accountRolesModule.AccountRolesRepo,
			hasher,
		),

		// ================= GETME =================
		GetMe: userucs.NewGetMeUC(
			userModule.UserQueryService,
		),

		// ================= GETALLEMPLOYESS ACCS WITH ROLES =================
		GetAllAccountsWithRolesUC: userucs.NewGetAllAccountsWithRolesUC(
			userModule.UserQueryService,
			useraccountpolicies.NewUserAccountCrudPolicy(),
		),

		// ================= EMPLOYEE ACCOUNT =================
		CreateEmployeeAcc: userucs.NewCreateEmployeeAccountUC(
			tx,
			genericcruduc.NewCRUDUseCase(userModule.UserRepo, userModule.AccountPolicies.CRUD),
			genericcruduc.NewCRUDUseCase(profileModule.ProfileRepo, profileModule.ProfilePolicies.CRUD),
			genericcruduc.NewCRUDUseCase(accountRolesModule.AccountRolesRepo, accountRolesModule.AccountRolePolicies.CRUD),
		),
	}
}
