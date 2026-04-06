package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StatusRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// statusRepo := gormentityrepos.NewGormStatusRepository(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Status,
	// 	dto.StatusCreateUpdateDTO,
	// 	dto.StatusResponseDTO,
	// ](status_crud_uc)

	// genericrouter.RegisterCRUDRoutes(r, "statuses", authMiddleware, tenantMiddleware, genericHandler)
}
