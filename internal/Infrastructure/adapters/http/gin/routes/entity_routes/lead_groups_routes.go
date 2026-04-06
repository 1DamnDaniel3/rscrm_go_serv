package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LeadGroupsRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// lead_groups_repo := gormentityrepos.NewGormLeadGroupsRepo(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.LeadGroup,
	// 	dto.LeadGroupCreateUpdateDTO,
	// 	dto.LeadGroupResponseDTO,
	// ](lead_groups_crudUC)

	// genericrouter.RegisterCRUDRoutes(r, "lead_groups", authMiddleware, tenantMiddleware, genericHandler)
}
