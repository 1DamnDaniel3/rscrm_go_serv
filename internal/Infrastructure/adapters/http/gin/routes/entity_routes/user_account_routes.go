package entityroutes

import (
	accountrolespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/account_roles_policies"
	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	userprofilepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_profile_policies"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/transactions"
	userhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user_handlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserAccountRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	hasher *bcrypt.BcryptHasher,
	tx services.Transaction,
	JwtSigner ports.JWTservice,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware) {
	// ==/== repo
	userRepo := gormentityrepos.NewGormUserAccountRepo(db, hasher)
	schoolRepo := gormentityrepos.NewGormSchoolRepo(db)
	profileRepo := gormentityrepos.NewGormUserProfileRepo(db)
	accRolesRepo := gormentityrepos.NewGormAccountRolesRepo(db)
	rolesRepo := gormentityrepos.NewGormRolesRepo(db)

	// ==/== policies
	crudPolicy := useraccountpolicies.NewUserAccountCrudPolicy()
	userAccountPolicies := useraccountpolicies.NewUserAccountPolicies(crudPolicy)

	// other policies
	profileCrudPolicy := userprofilepolicies.NewUserProfieCrudPolicy()
	accRolesCrudPolicy := accountrolespolicies.NewAccountRolesCrudPolicy()

	// ==/== generic user

	userAccountCrudUC := genericcruduc.NewCRUDUseCase(userRepo, userAccountPolicies.CRUD)

	genericUserHandler := generichandler.NewGenericHandler[
		entities.UserAccount,
		dto.UserAccountCreateDTO,
		dto.UserAccountResponseDTO,
	](userAccountCrudUC)

	// ==/== login
	// userAccountCrudUC := genericcruduc.NewCRUDUseCase(userRepo, userAccountPolicies.CRUD)
	loginUseCase := userUCs.NewLoginUseCase(hasher, userRepo, accRolesRepo, rolesRepo, JwtSigner)
	loginHandler := userhandlers.NewLoginHandler(loginUseCase)

	// ==/== register
	registerUC := userUCs.NewRegisterUseCase(tx, userRepo, profileRepo, schoolRepo, accRolesRepo, hasher)
	registerHandler := transactions.NewRegisterHandler(registerUC)

	authCheckHandler := userhandlers.NewAuthCheckHandler(JwtSigner)

	// ==/== create employee account
	accountCrudUC := genericcruduc.NewCRUDUseCase(userRepo, userAccountPolicies.CRUD)
	profileCrudUC := genericcruduc.NewCRUDUseCase(profileRepo, profileCrudPolicy)
	accountRolesCrudUC := genericcruduc.NewCRUDUseCase(accRolesRepo, accRolesCrudPolicy)

	createEmpAccUC := userUCs.NewCreateEmployeeAccountUC(
		tx,
		accountCrudUC,
		profileCrudUC,
		accountRolesCrudUC,
	)
	CreateEmpAccHandler := userhandlers.NewCreateEmployeeAccountHandler(createEmpAccUC)

	// ==/== routes

	protected := genericrouter.RegisterCRUDRoutes(r, "user_accounts", authMiddleware, tenantMiddleware, genericUserHandler)

	r.GET("/auth/check", authCheckHandler.CheckAuth)
	r.POST("/ownerschool/register", registerHandler.Register)
	r.POST("/user_accounts/login", loginHandler.Login)
	r.GET("/user_accounts/logout", loginHandler.Logout)

	protected.POST("/user_accounts/createemployee", CreateEmpAccHandler.CreateEmployeeAccountHandler)

	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)

	// r.POST("/user_accounts/create")
}
