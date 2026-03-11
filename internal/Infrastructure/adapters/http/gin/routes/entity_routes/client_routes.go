package entityroutes

import (
	clientucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	clienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/clientHandlers"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ClientRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	clientRepo := gormentityrepos.NewGormClientRepo(db)
	clientQueryService := gormentityrepos.NewGormClientQueryService(db)
	clientGroupRepo := gormentityrepos.NewGormClientGroupRepo(db)

	// Get Grouped
	groupedUC := clientucs.NewGroupedClientsUC(clientRepo)
	groupedHandler := clienthandlers.NewGroupedClientsHandler(groupedUC)

	// Create Grouped
	createGroupedUC := clientucs.NewCreateGroupedClientUC(tx, clientGroupRepo, clientRepo)
	createGropedHandler := clienthandlers.NewCreateGroupedClientHandler(createGroupedUC)

	// Search
	searchHandler := clienthandlers.NewClientSearchHandler(clientQueryService)

	// Client Groups
	clientGroupsUC := clientucs.NewGetClientGroupUC(clientQueryService)
	clientGroupsHandler := clienthandlers.NewGetClientGroupsHandler(clientGroupsUC)

	// Client Students
	clientStudentsUC := clientucs.NewGetClientStudentsUC(clientQueryService)
	clientStudentsHandler := clienthandlers.NewGetClientStudentsHandler(clientStudentsUC)

	// Generic
	genericHandler := genericHandler.NewGenericHandler[
		entities.Client,
		dto.ClientCreateUpdateDTO,
		dto.ClientResponseDTO,
	](clientRepo)

	protected := genericrouter.RegisterCRUDRoutes(r, "clients", authMiddleware, tenantMiddleware, genericHandler)
	protected.POST("clients/groupedclients", groupedHandler.GetGroupedClients)
	protected.POST("clients/createandgroup", createGropedHandler.CreateGroupedClient)
	protected.GET("clients/search", searchHandler.Search)
	protected.GET("clients/:id/groups", clientGroupsHandler.GetGroups)
	protected.GET("clients/:id/students", clientStudentsHandler.GetClientStudents)

}
