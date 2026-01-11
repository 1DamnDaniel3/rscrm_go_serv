package main

import (
	"log"
	"os"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/config"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

// генерация swagger
// swag init -g cmd/server/main.go

func main() {
	_ = godotenv.Load()
	PORT := os.Getenv("PORT")

	router := gin.Default()

	if os.Getenv("ENV") == "dev" {
		router.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
			AllowCredentials: true,
		}))
	}

	db_cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config :( because: %v", err)
	}

	db, err := adapters.ConnectDB(db_cfg)
	if err != nil {
		log.Fatalf("failed to connectDB :( because: %v", err)
	}
	routes.SetupRoutes(router, db)

	// migrations

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDB.Close()
	if err := goose.Up(sqlDB, "./migrations"); err != nil { // migrations up
		log.Fatal(err)
	}

	router.Run(":" + PORT)
}
