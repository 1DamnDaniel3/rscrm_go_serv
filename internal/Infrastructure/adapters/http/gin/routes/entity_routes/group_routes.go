package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GroupRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// groupRepo := gormentityrepos.NewGormGroupRepository(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Group,
	// 	dto.GroupCreateUpdateDTO,
	// 	dto.GroupResponseDTO,
	// ](groupCrudUC)

	// genericrouter.RegisterCRUDRoutes(r, "groups", authMiddleware, tenantMiddleware, genericHandler)
}
