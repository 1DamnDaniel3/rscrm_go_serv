package entityroutes

import (
	lessonpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lesson_policies"
	lessonsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs/lessonShedulesUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	lessonhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/lessonHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func LessonRoutes(
	r *gin.RouterGroup,
	db *gorm.DB,
	tx services.Transaction,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {
	// ==/== repo
	lessonRepo := gormentityrepos.NewGormLessonRepo(db)
	lessonQueryService := gormentityrepos.NewGormLessonQueryService(db)
	scheduleRepo := gormentityrepos.NewGormScheduleRepo(db)

	// ==/== policies
	crudPolicy := lessonpolicies.NewLessonCrudPolicy()
	lessonPolicies := lessonpolicies.NewLessonPolicies(crudPolicy)

	lessonCrudUC := genericcruduc.NewCRUDUseCase(lessonRepo, lessonPolicies.CRUD)

	genericHandler := generichandler.NewGenericHandler[
		entities.Lesson,
		dto.LessonCreateUpdateDTO,
		dto.LessonResponseDTO,
	](lessonCrudUC)

	// generate
	generateLessonUC := lessonshedulesucs.NewCreateLessonsFromShceduleUC(lessonRepo, scheduleRepo)

	// cleanUp
	cleanupUC := lessonsucs.NewCleanupOldLessonsUC(lessonQueryService)

	// fetch lessons
	fetchLessonsUC := lessonsucs.NewGetLessonsUC(lessonRepo, generateLessonUC, cleanupUC)
	fetchLessonsHandler := lessonhandlers.NewGetLessonsHandler(fetchLessonsUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "lessons", authMiddleware, tenantMiddleware, genericHandler)

	// protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)
	protected.GET("/lessons/fetch", fetchLessonsHandler.GetLessons)
}
