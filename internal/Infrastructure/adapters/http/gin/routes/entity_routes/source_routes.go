package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SourceRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// sourceRepo := gormentityrepos.NewGormSourceRepository(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Source,
	// 	dto.SourceCreateUpdateDTO,
	// 	dto.SourceResponseDTO,
	// ](sourceCrudUC)

	// genericrouter.RegisterCRUDRoutes(r, "sources", authMiddleware, tenantMiddleware, genericHandler)
}
