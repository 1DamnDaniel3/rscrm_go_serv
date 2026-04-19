package routes

import (
	"os"
	"time"

	_ "github.com/1DamnDaniel3/rscrm_go_serv/docs"
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	entityroutes "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/jwt"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders"
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

	// ==========================
	// |		  APP		     |
	// ==========================

	app := infrastructure.NewAppContainer(db, hasher, JWTSigner, tx)
	usecases := builders.NewUseCases(app)

	// == routes ==
	entityroutes.UserAccountRoutes(api, app, usecases.UserAccount, authMiddleware, tenandMiddleware)
	entityroutes.UserProfileRoutes(api, app, usecases.UserProfile, authMiddleware, tenandMiddleware)
	entityroutes.SchoolRoutes(api, app, usecases.School, authMiddleware, tenandMiddleware)
	entityroutes.LeadRoutes(api, app, usecases.Lead, usecases.LeadGroup, authMiddleware, tenandMiddleware)
	entityroutes.StudentRoutes(api, app, usecases.Student, usecases.StudentGroup, authMiddleware, tenandMiddleware)
	entityroutes.ClientRoutes(api, app, usecases.Client, authMiddleware, tenandMiddleware)
	entityroutes.GroupRoutes(api, app, usecases.Group, authMiddleware, tenandMiddleware)
	entityroutes.StatusRoutes(api, app, usecases.Status, authMiddleware, tenandMiddleware)
	entityroutes.SourceRoutes(api, app, usecases.Source, authMiddleware, tenandMiddleware)
	entityroutes.ScheduleRoutes(api, app, usecases.Schedule, authMiddleware, tenandMiddleware)
	entityroutes.LessonRoutes(api, app, usecases.Lesson, authMiddleware, tenandMiddleware)

	// == related tables routes
	entityroutes.StudentClientRoutes(api, app, usecases.StudentClients, authMiddleware, tenandMiddleware)

	entityroutes.LeadGroupsRoutes(api, app, usecases.LeadGroup, authMiddleware, tenandMiddleware)
}
