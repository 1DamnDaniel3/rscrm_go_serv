package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	accountrolesbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/account_roles_builders"
	"github.com/gin-gonic/gin"
)

func AccountRolesRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	accountRolesUCs *accountrolesbuilders.AccountRolesUseCases,
) {

	// ================= Handlers =================
	accRolesHandlers := accountrolesbuilders.NewAccountRolesHandlerBuilder(accountRolesUCs)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	protected.POST("/user_accounts/assignroles", accRolesHandlers.NotGenericCrudHandler.AssingRoles)
	protected.DELETE("/user_accounts/:acc_id/roles/:role_id", accRolesHandlers.NotGenericCrudHandler.RemoveRoles)
}
