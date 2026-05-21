package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	transactionsbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/transactions_builders"
	"github.com/gin-gonic/gin"
)

func TransactionRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	transactionsUCs *transactionsbuilders.TransactionsUseCases,
) {
	// ================= Handlers =================
	handlers := transactionsbuilders.NewSubscriptionHandlerBuilder(
		transactionsUCs,
	)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	protected.GET("/transactions", handlers.CRUDHandler.GetAll)
}
