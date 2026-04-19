package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	clientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/client_builders"
	"github.com/gin-gonic/gin"
)

func ClientRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	clientUCs *clientbuilders.ClientUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	clientHandlers := clientbuilders.NewClientHandlerBuilder(clientUCs, app.ClientModule)

	// ================= Routes =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"clients",
		authMiddleware,
		tenantMiddleware,
		clientHandlers.CRUDHandler,
	)

	protected.POST("/clients/groupedclients", clientHandlers.GroupedClientsHandler.GetGroupedClients)
	protected.POST("/clients/createandgroup", clientHandlers.CreateGroupedHandler.CreateGroupedClient)
	protected.GET("/clients/search", clientHandlers.SearchHandler.Search)
	protected.GET("/clients/:id/groups", clientHandlers.GetClientGroups.GetGroups)
	protected.GET("/clients/:id/students", clientHandlers.GetClientStudents.GetClientStudents)
}
