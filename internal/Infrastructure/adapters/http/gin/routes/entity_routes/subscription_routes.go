package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	subscriptionbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/subscription_builders"
	"github.com/gin-gonic/gin"
)

func SubscriptionRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	subscriptionUCs *subscriptionbuilders.SubscriptionUseCases,
) {
	// ================= Handlers =================
	handlers := subscriptionbuilders.NewSubscriptionHandlerBuilder(
		subscriptionUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"subscriptions",
		app.AuthMiddleware,
		app.TenantMiddleware,
		handlers.CRUDHandler,
	)
}
