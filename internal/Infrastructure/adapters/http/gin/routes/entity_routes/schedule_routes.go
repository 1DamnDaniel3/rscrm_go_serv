package entityroutes

import (
	schedulepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/schedule_policies"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ScheduleRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	scheduleRepo := gormentityrepos.NewGormScheduleRepo(db)

	// ==/== policies

	crudPolicy := schedulepolicies.NewScheduleCrudPolicy()
	schedulePolicies := schedulepolicies.NewSchedulePolicies(crudPolicy)

	scheduleCrudUC := genericcruduc.NewCRUDUseCase(scheduleRepo, schedulePolicies.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.Schedule,
		dto.ScheduleCreateUpdateDTO,
		dto.ScheduleResponseDTO,
	](scheduleCrudUC)

	genericrouter.RegisterCRUDRoutes(r, "schedules", authMiddleware, tenantMiddleware, genericHandler)
}
