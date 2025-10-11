package routes

import (
	_ "github.com/1DamnDaniel3/rscrm_go_serv/docs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/user"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
	adapters "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	entityroutes "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes"
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

	// == passwordHasher ==
	hasher := bcrypt.NewBcryptHasher(12)

	// == GORM_transaction ==
	tx := adapters.NewGormTransaction(db)

	// == Repositories GENERIC ==
	profileRepo := generic.NewGormRepository[entities.UserProfile](db)
	schoolRepo := generic.NewGormRepository[entities.School](db)

	// == Repositories Entities ==
	userRepo := adapters.NewGormUserAccountRepo(db, hasher)

	// == usecases ==
	registerUC := user.NewRegisterUseCase(tx, userRepo, profileRepo, schoolRepo, hasher)

	// == routes ==
	entityroutes.UserAccountRoutes(api, userRepo)
	entityroutes.RegisterRouter(api, registerUC)
}
