package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	userUC "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/transactions"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
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
	// ==/== repos
	userRepo := gormentityrepos.NewGormUserAccountRepo(db, hasher)
	schoolRepo := gormentityrepos.NewGormSchoolRepo(db)
	profileRepo := gormentityrepos.NewGormUserProfileRepo(db)

	// ==/== policies
	// crudPolicy := useraccountpolicies.NewUserAccountCrudPolicy()
	// userAccountPolicies := useraccountpolicies.NewUserAccountPolicies(crudPolicy)

	accRolesRepo := gormentityrepos.NewGormAccountRolesRepo(db)
	rolesRepo := gormentityrepos.NewGormRolesRepo(db)
	// ==/== usecases
	// userAccountCrudUC := genericcruduc.NewCRUDUseCase(userRepo, userAccountPolicies.CRUD)
	loginUseCase := userUC.NewLoginUseCase(hasher, userRepo, accRolesRepo, rolesRepo, JwtSigner)
	registerUC := userUC.NewRegisterUseCase(tx, userRepo, profileRepo, schoolRepo, accRolesRepo, hasher)

	// ==/== handlers
	// genericUserHandler := generichandler.NewGenericHandler[
	// 	entities.UserAccount,
	// 	dto.UserAccountCreateDTO,
	// 	dto.UserAccountResponseDTO,
	// ](userAccountCrudUC)

	loginHandler := user.NewLoginHandler(loginUseCase)
	registerHandler := transactions.NewRegisterHandler(registerUC)
	authCheckHandler := user.NewAuthCheckHandler(JwtSigner)

	// ==/== routes

	// genericRoute.RegisterCRUDRoutes(r, "user_accounts", authMiddleware, tenantMiddleware, genericUserHandler)

	r.GET("/auth/check", authCheckHandler.CheckAuth)
	r.POST("/ownerschool/register", registerHandler.Register)
	r.POST("/user_accounts/login", loginHandler.Login)
	r.GET("/user_accounts/logout", loginHandler.Logout)

	// r.POST("/user_accounts/create")
}
