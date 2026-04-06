package entityroutes

import (
	studentclientUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs/student_clientUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"

	studentclienthandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/studentClientHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
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

	// stud_cliHandler := genericHandler.NewGenericHandler[
	// 	entities.StudentClient,
	// 	dto.StudentClientCreateUpdateDTO,
	// 	dto.StudentClientResponseDTO,
	// ](stud_cli_crud_UC)

	// create and get BO
	CreateAndGetBOUC := studentclientUCs.NewCreateRelationUC(tx, stud_cliRepo, clientRepo)
	CreateRelationHandler := studentclienthandlers.NewCreateStudentClientRelHandler(CreateAndGetBOUC)

	// protected := genericrouter.RegisterCRUDRoutes(r, "student-clients", authMiddleware, tenantMiddleware, stud_cliHandler)
	protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)
	protected.POST("/student-clients/createandgetBO", CreateRelationHandler.CreateRel)
}
