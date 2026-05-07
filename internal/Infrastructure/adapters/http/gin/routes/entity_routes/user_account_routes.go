package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	useraccountbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_accounts_builders"
	"github.com/gin-gonic/gin"
)

func UserAccountRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *useraccountbuilders.UserUseCases,
) {

	// ================= handlers =================
	handlers := useraccountbuilders.NewUserAccountHandlerBuilder(useCases, app)

	// ================= CRUD =================
	protected := genericrouter.GetProtectedRouterGroup(r, app.AuthMiddleware, app.TenantMiddleware)
	// protected := genericrouter.RegisterCRUDRoutes(
	// 	r,
	// 	"user_accounts",
	// 	app.AuthMiddleware,
	// 	app.TenantMiddleware,
	// 	handlers.CRUDHandler,
	// )
	// +++++++++++++++++++++++++++++++++ Достать все акканту + их роли
	protected.POST("/user_accounts", handlers.CRUDHandler.Create)
	protected.PATCH("/user_accounts/:id", handlers.CRUDHandler.Update)

	// ================= auth =================
	r.POST("/ownerschool/register", handlers.RegisterHandler.Register)
	r.POST("/user_accounts/login", handlers.LoginHandler.Login)
	r.GET("/user_accounts/logout", handlers.LoginHandler.Logout)
	r.GET("/auth/check", handlers.AuthCheckHandler.CheckAuth)

	// -=== private
	protected.GET("/me", handlers.GetMeHandler.GetMe)

	// ================= employee =================
	protected.GET("/user_accounts/allwithroles", handlers.GetAllAccountsWithRolesHandler.GetAccountsWithRoles)
	protected.POST("/user_accounts/createemployee", handlers.CreateEmployeeHandler.CreateEmployeeAccountHandler)
}
