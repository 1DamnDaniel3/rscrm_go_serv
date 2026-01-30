package genericrouter

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/gin-gonic/gin"
)

// Auto generate unique routes for basic CRUD operations.
// r for RouterGroup like `/api/somethings/`
// Prefix its REST oriented entity name in route like `/cars/`.
// gin handler with *gin.Context
func RegisterCRUDRoutes[T any, C any, R any](r *gin.RouterGroup,
	prefix string,
	authMiddleware *middleware.AuthMiddleware,
	handler *generic.GenericHandler[T, C, R],
) {

	routeGroup := r
	if authMiddleware != nil {
		routeGroup = r.Group("")
		routeGroup.Use(authMiddleware.TryAuth())
	}

	routeGroup.POST("/"+prefix+"/create", handler.Create)
	routeGroup.PATCH("/"+prefix+"/update/:id", handler.Update)
	routeGroup.POST("/"+prefix+"/getallwhere", handler.GetAllWhere)
	routeGroup.GET("/"+prefix+"/getone/:id", handler.GetByID)
	routeGroup.GET("/"+prefix+"/getall", handler.GetAll)
	routeGroup.DELETE("/"+prefix+"/delete/:id", handler.Delete)

}
