package routes

import (
	_ "github.com/1DamnDaniel3/rscrm_go_serv/docs"
	entityroutes "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/routes/entity_routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFile("/swagger-crud.json", "./docs/swagger-crud.json")
	r.GET("/swagger-crud/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger-crud.json")))

	entityroutes.UserAccountRoutes(api, db)
}
