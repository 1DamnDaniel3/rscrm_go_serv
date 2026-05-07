package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	userprofilebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_profile_builders"

	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *userprofilebuilders.UserProfileUseCases,
) {

	// ================= handlers =================
	handlers := userprofilebuilders.NewUserProfileHandlerBuilder(useCases)

	// ================= protected =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	// ================= CRUD =================
	protected.PATCH("user_accounts/:id/profiles", handlers.CRUDHandler.Update)

	// ================= routes =================
	protected.GET("/user_accounts/:id/profiles", handlers.CRUDHandler.GetByID)
	protected.GET("/user_profiles", handlers.CRUDHandler.GetAll)
	protected.POST("/user_profiles/getallwhere", handlers.CRUDHandler.GetAllWhere)
}
