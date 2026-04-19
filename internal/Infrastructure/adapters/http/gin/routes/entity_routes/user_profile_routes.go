package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	userprofilebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_profile_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func UserProfileRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	useCases *userprofilebuilders.UserProfileUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= handlers =================
	handlers := userprofilebuilders.NewUserProfileHandlerBuilder(useCases)

	// ================= protected =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		authMiddleware,
		tenantMiddleware,
	)

	// ================= routes =================
	protected.GET(
		"/user_account/:id/profile",
		handlers.CRUDHandler.GetByID,
	)
}
