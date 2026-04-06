package entityroutes

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SchoolRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// schoolRepo := gormentityrepos.NewGormSchoolRepo(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.School,
	// 	dto.SchoolCreateUpdateDTO,
	// 	dto.SchoolResponseDTO,
	// ](school_crud_UC)

	// genericrouter.RegisterCRUDRoutes(r, "schools", authMiddleware, tenantMiddleware, genericHandler)
}
