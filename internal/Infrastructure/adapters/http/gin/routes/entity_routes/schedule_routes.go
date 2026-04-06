package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
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

	// scheduleRepo := gormentityrepos.NewGormScheduleRepo(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Schedule,
	// 	dto.ScheduleCreateUpdateDTO,
	// 	dto.ScheduleResponseDTO,
	// ](scheduleCrudUC)

	// genericrouter.RegisterCRUDRoutes(r, "schedules", authMiddleware, tenantMiddleware, genericHandler)

}
