package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	dancestylebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/dance_style_builders"
	"github.com/gin-gonic/gin"
)

func DanceStyleRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	danceStyleUCs *dancestylebuilders.DanceStyleUseCases,
) {

	// ================= Handlers =================
	danceStyleHandlers := dancestylebuilders.NewClientHandlerBuilder(danceStyleUCs, app.DanceStyleModule)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	protected.GET("/dance_styles/getall", danceStyleHandlers.CRUDHandler.GetAll)
}
