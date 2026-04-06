package entityroutes

import (
	lessonsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs/lessonShedulesUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	lessonhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/lessonHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
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
	// repo
	lessonRepo := gormentityrepos.NewGormLessonRepo(db)
	lessonQueryService := gormentityrepos.NewGormLessonQueryService(db)
	scheduleRepo := gormentityrepos.NewGormScheduleRepo(db)

	// genericHandler := genericHandler.NewGenericHandler[
	// 	entities.Lesson,
	// 	dto.LessonCreateUpdateDTO,
	// 	dto.LessonResponseDTO,
	// ](lessonCrudUC)

	// generate
	generateLessonUC := lessonshedulesucs.NewCreateLessonsFromShceduleUC(lessonRepo, scheduleRepo)

	// cleanUp
	cleanupUC := lessonsucs.NewCleanupOldLessonsUC(lessonQueryService)

	// fetch lessons
	fetchLessonsUC := lessonsucs.NewGetLessonsUC(lessonRepo, generateLessonUC, cleanupUC)
	fetchLessonsHandler := lessonhandlers.NewGetLessonsHandler(fetchLessonsUC)

	// protected := genericrouter.RegisterCRUDRoutes(r, "lessons", authMiddleware, tenantMiddleware, genericHandler)

	protected := genericrouter.GetProtectedRouterGroup(r, authMiddleware, tenantMiddleware)
	protected.GET("/lessons/fetch", fetchLessonsHandler.GetLessons)
}
