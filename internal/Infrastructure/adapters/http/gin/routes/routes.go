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

	// ===========================
	// |          APP            |
	// ===========================

	app := infrastructure.NewAppContainer(
		db, hasher,
		JWTSigner, tx,
		authMiddleware, tenandMiddleware,
	)
	usecases := builders.NewUseCases(app)

	// == routes ==
	entityroutes.AccountRolesRoutes(api, app, usecases.AccountRoles)
	entityroutes.SchoolRoutes(api, app, usecases.School)
	entityroutes.UserAccountRoutes(api, app, usecases.UserAccount)
	entityroutes.UserProfileRoutes(api, app, usecases.UserProfile)

	entityroutes.LeadRoutes(api, app, usecases.Lead, usecases.LeadGroup)
	entityroutes.ClientRoutes(api, app, usecases.Client)
	entityroutes.StudentRoutes(api, app, usecases.Student, usecases.StudentGroup)

	entityroutes.GroupRoutes(api, app, usecases.Group)

	entityroutes.StatusRoutes(api, app, usecases.Status)
	entityroutes.SourceRoutes(api, app, usecases.Source)

	entityroutes.ScheduleRoutes(api, app, usecases.Schedule)
	entityroutes.LessonRoutes(api, app, usecases.Lesson)

	entityroutes.SubscriptionRoutes(api, app, usecases.Subscription)

	// == related tables routes
	entityroutes.StudentClientRoutes(api, app, usecases.StudentClients, authMiddleware, tenandMiddleware)

	entityroutes.LeadGroupsRoutes(api, app, usecases.LeadGroup, authMiddleware, tenandMiddleware)
}
