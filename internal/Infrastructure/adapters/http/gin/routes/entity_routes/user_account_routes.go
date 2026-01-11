package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	userUC "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/transactions"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user"
	genericRoute "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserAccountRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	hasher *bcrypt.BcryptHasher,
	tx services.Transaction,
	JwtSigner ports.JWTservice) {
	// ==/== repos
	userRepo := adapters.NewGormUserAccountRepo(db, hasher)
	schoolRepo := adapters.NewGormSchoolRepo(db)
	profileRepo := adapters.NewGormUserProfileRepo(db)

	accRolesRepo := adapters.NewGormAccountRolesRepo(db)
	rolesRepo := adapters.NewGormRolesRepo(db)
	// ==/== usecases
	LoginUseCase := userUC.NewLoginUseCase(hasher, userRepo, accRolesRepo, rolesRepo, JwtSigner)
	registerUC := userUC.NewRegisterUseCase(tx, userRepo, profileRepo, schoolRepo, accRolesRepo, hasher)

	// ==/== handlers
	genericUserHandler := genericHandler.NewGenericHandler[
		entities.UserAccount,
		dto.UserAccountCreateDTO,
		dto.UserAccountResponseDTO,
	](userRepo)
	loginHandler := user.NewLoginHandler(LoginUseCase)
	registerHandler := transactions.NewRegisterHandler(registerUC)
	authCheckHandler := user.NewAuthCheckHandler(JwtSigner)

	// ==/== routes

	genericRoute.RegisterCRUDRoutes(r, "user_accounts", genericUserHandler)

	r.GET("/auth/check", authCheckHandler.CheckAuth)
	r.POST("/ownerschool/register", registerHandler.Register)
	r.POST("/user_accounts/login", loginHandler.Login)

	// r.POST("/user_accounts/create")
}
