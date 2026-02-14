package entityroutes

import (
	clientucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	clienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/clientHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
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
	clientGroupRepo := gormentityrepos.NewGormClientGroupRepo(db)

	// Get Grouped
	groupedUC := clientucs.NewGroupedClientsUC(clientRepo)
	groupedHandler := clienthandlers.NewGroupedClientsHandler(groupedUC)

	// Create Grouped
	createGroupedUC := clientucs.NewCreateGroupedClientUC(tx, clientGroupRepo, clientRepo)
	createGropedHandler := clienthandlers.NewCreateGroupedClientHandler(createGroupedUC)

	genericHandler := generic.NewGenericHandler[
		entities.Client,
		dto.ClientCreateUpdateDTO,
		dto.ClientResponseDTO,
	](clientRepo)

	protected := genericrouter.RegisterCRUDRoutes(r, "clients", authMiddleware, tenantMiddleware, genericHandler)
	protected.POST("clients/groupedclients", groupedHandler.GetGroupedClients)
	protected.POST("clients/createandgroup", createGropedHandler.CreateGroupedClient)

}
