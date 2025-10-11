package generichandler

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/generic"
	"github.com/gin-gonic/gin"
)

func RegisterCRUDRoutes[T any, C any, R any](r *gin.RouterGroup, prefix string, handler *generic.GenericHandler[T, C, R]) {
	r.POST("/"+prefix+"/register", handler.Create)
	r.POST("/"+prefix+"/update/:id", handler.Update)
	r.POST("/"+prefix+"/getallwhere", handler.GetAllWhere)
	r.GET("/"+prefix+"/getone/:id", handler.GetByID)
	r.GET("/"+prefix+"/getall", handler.GetAll)
	r.DELETE("/"+prefix+"/delete/:id", handler.Delete)
}
