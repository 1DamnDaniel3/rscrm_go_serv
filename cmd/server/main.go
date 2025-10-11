package main

import (
	"log"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/config"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db_cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := adapters.ConnectDB(db_cfg)
	if err != nil {
		log.Fatalf("failed to connectDB: %v", err)
	}
	routes.SetupRoutes(router, db)

	router.Run(":3001")
}
