package entityroutes

import (
	studentclientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_client_policies"
	studentclientUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/studentUCs/student_clientUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
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
	// ==/== repo
	stud_cliRepo := gormentityrepos.NewGormStudentClientsRepo(db)
	clientRepo := gormentityrepos.NewGormClientRepo(db)

	// ==/== policies
	crudPolicy := studentclientpolicies.NewStudentClientCrudPolicy()
	studCliPolicies := studentclientpolicies.NewStudentClientPolicies(crudPolicy)

	studCliCrudUC := genericcruduc.NewCRUDUseCase(stud_cliRepo, studCliPolicies.CRUD)

	stud_cliHandler := generichandler.NewGenericHandler[
		entities.StudentClient,
		dto.StudentClientCreateUpdateDTO,
		dto.StudentClientResponseDTO,
	](studCliCrudUC)

	// create and get BO
	CreateAndGetBOUC := studentclientUCs.NewCreateRelationUC(tx, stud_cliRepo, clientRepo)
	CreateRelationHandler := studentclienthandlers.NewCreateStudentClientRelHandler(CreateAndGetBOUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "student-clients", authMiddleware, tenantMiddleware, stud_cliHandler)
	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)
	protected.POST("/student-clients/createandgetBO", CreateRelationHandler.CreateRel)
}
