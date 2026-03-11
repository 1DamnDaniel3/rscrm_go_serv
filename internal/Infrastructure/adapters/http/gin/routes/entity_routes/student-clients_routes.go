package entityroutes

import (
	studentclientsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentClientsUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"

	studentclienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentClientHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StudentClientRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	stud_cliRepo := gormentityrepos.NewGormStudentClientsRepo(db)
	clientRepo := gormentityrepos.NewGormClientRepo(db)

	stud_cliHandler := genericHandler.NewGenericHandler[
		entities.StudentClient,
		dto.StudentClientCreateUpdateDTO,
		dto.StudentClientResponseDTO,
	](stud_cliRepo)

	// create and get BO
	CreateAndGetBOUC := studentclientsucs.NewCreateRelationUC(tx, stud_cliRepo, clientRepo)
	CreateRelationHandler := studentclienthandlers.NewCreateStudentClientRelHandler(CreateAndGetBOUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "student-clients", authMiddleware, tenantMiddleware, stud_cliHandler)
	protected.POST("/student-clients/createandgetBO", CreateRelationHandler.CreateRel)
}
