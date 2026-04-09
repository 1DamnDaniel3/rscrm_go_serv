package entityroutes

import (
	clientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_policies"
	clientgroupsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/clientUCs/client_groupUCs"
	clientstudentsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/clientUCs/client_studentUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	clienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/clientHandlers"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
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
	// repo
	clientRepo := gormentityrepos.NewGormClientRepo(db)

	// policies
	crudPolicy := clientpolicies.NewClientCrudPolicy()
	clientPolicies := clientpolicies.NewClientPolicies(crudPolicy)

	clientCrudUC := genericcruduc.NewCRUDUseCase(clientRepo, clientPolicies.CRUD)

	clientQueryService := gormentityrepos.NewGormClientQueryService(db)
	clientGroupRepo := gormentityrepos.NewGormClientGroupRepo(db)

	// Get Grouped
	groupedUC := clientgroupsUCs.NewGroupedClientsUC(clientRepo)
	groupedHandler := clienthandlers.NewGroupedClientsHandler(groupedUC)

	// Create Grouped
	createGroupedUC := clientgroupsUCs.NewCreateGroupedClientUC(tx, clientGroupRepo, clientRepo)
	createGropedHandler := clienthandlers.NewCreateGroupedClientHandler(createGroupedUC)

	// Search
	searchHandler := clienthandlers.NewClientSearchHandler(clientQueryService)

	// Client Groups
	clientGroupsUC := clientgroupsUCs.NewGetClientGroupUC(clientQueryService)
	clientGroupsHandler := clienthandlers.NewGetClientGroupsHandler(clientGroupsUC)

	// Client Students
	clientStudentsUC := clientstudentsUCs.NewGetClientStudentsUC(clientQueryService)
	clientStudentsHandler := clienthandlers.NewGetClientStudentsHandler(clientStudentsUC)

	// Generic
	genericHandler := generichandler.NewGenericHandler[
		entities.Client,
		dto.ClientCreateUpdateDTO,
		dto.ClientResponseDTO,
	](clientCrudUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "clients", authMiddleware, tenantMiddleware, genericHandler)
	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)

	protected.POST("/clients/groupedclients", groupedHandler.GetGroupedClients)
	protected.POST("/clients/createandgroup", createGropedHandler.CreateGroupedClient)
	protected.GET("/clients/search", searchHandler.Search)
	protected.GET("/clients/:id/groups", clientGroupsHandler.GetGroups)
	protected.GET("/clients/:id/students", clientStudentsHandler.GetClientStudents)

}
