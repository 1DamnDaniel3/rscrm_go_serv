package entityroutes

import (
	lessonsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/lessonsUCs"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/lessonsUCs/lessonShedulesUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	genericHandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
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
	// repo
	lessonRepo := gormentityrepos.NewGormLessonRepo(db)
	lessonQueryService := gormentityrepos.NewGormLessonQueryService(db)
	scheduleRepo := gormentityrepos.NewGormScheduleRepo(db)

	genericHandler := genericHandler.NewGenericHandler[
		entities.Lesson,
		dto.LessonCreateUpdateDTO,
		dto.LessonResponseDTO,
	](lessonRepo)

	// generate
	generateLessonUC := lessonshedulesucs.NewCreateLessonsFromShceduleUC(lessonRepo, scheduleRepo)

	// cleanUp
	cleanupUC := lessonsucs.NewCleanupOldLessonsUC(lessonQueryService)

	// fetch lessons
	fetchLessonsUC := lessonsucs.NewGetLessonsUC(lessonRepo, generateLessonUC, cleanupUC)
	fetchLessonsHandler := lessonhandlers.NewGetLessonsHandler(fetchLessonsUC)

	protected := genericrouter.RegisterCRUDRoutes(r, "lessons", authMiddleware, tenantMiddleware, genericHandler)

	protected.GET("/lessons/fetch", fetchLessonsHandler.GetLessons)
}
