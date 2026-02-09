package routes

import (
	"os"
	"time"

	_ "github.com/1DamnDaniel3/rscrm_go_serv/docs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	entityroutes "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/jwt"
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
	// ====================shared dependences

	// == passwordHasher ==
	hasher := bcrypt.NewBcryptHasher(12)

	// == GORM_transaction ==
	tx := adapters.NewGormTransaction(db)

	// == JWT ==
	secret := os.Getenv("JWT_SECRET")
	JWTSigner := jwt.NewJWTAdapter(secret, 5*time.Hour)
	authMiddleware := middleware.NewAuthMiddleware(JWTSigner)
	tenandMiddleware := middleware.NewTenandMiddleware()

	// == routes ==

	entityroutes.UserAccountRoutes(api, db, hasher, tx, JWTSigner, authMiddleware, tenandMiddleware)
	entityroutes.UserProfileRoutes(api, db, authMiddleware, tenandMiddleware)
	entityroutes.SchoolRoutes(api, db, authMiddleware, tenandMiddleware)
	entityroutes.LeadRoutes(api, db, tx, authMiddleware, tenandMiddleware)
	entityroutes.StudentRoutes(api, db, authMiddleware, tenandMiddleware)
	entityroutes.GroupRoutes(api, db, authMiddleware, tenandMiddleware)
	entityroutes.StatusRoutes(api, db, authMiddleware, tenandMiddleware)
	entityroutes.SourceRoutes(api, db, authMiddleware, tenandMiddleware)

	// == related tables routes

	entityroutes.LeadGroupsRoutes(api, db, authMiddleware, tenandMiddleware)
}
