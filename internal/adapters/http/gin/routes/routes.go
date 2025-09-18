package routes

import (
	entityroutes "github.com/1DamnDaniel3/rscrm_go_serv/internal/adapters/http/gin/routes/entity_routes"
	_ "github.com/1DamnDaniel3/rscrm_go_serv/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	entityroutes.UserAccountRoutes(api, db)
}
