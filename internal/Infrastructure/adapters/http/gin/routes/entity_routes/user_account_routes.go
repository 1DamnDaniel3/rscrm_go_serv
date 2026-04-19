package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	useraccountbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_accounts_builders"
	"github.com/gin-gonic/gin"
)

func UserAccountRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *useraccountbuilders.UserUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= handlers (локально) =================
	handlers := useraccountbuilders.NewUserAccountHandlerBuilder(useCases, app)

	// ================= CRUD =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"user_accounts",
		authMiddleware,
		tenantMiddleware,
		handlers.CRUDHandler,
	)

	// ================= public =================
	r.GET("/auth/check", handlers.AuthCheckHandler.CheckAuth)
	r.POST("/ownerschool/register", handlers.RegisterHandler.Register)
	r.POST("/user_accounts/login", handlers.LoginHandler.Login)
	r.GET("/user_accounts/logout", handlers.LoginHandler.Logout)

	// ================= employee =================
	protected.POST(
		"/user_accounts/createemployee",
		handlers.CreateEmployeeHandler.CreateEmployeeAccountHandler,
	)
}
