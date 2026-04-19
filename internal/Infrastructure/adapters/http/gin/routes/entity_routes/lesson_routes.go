package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	lessonbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lesson_builders"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func LessonRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	lessonUCs *lessonbuilders.LessonUseCases,
	authMiddleware *middleware.AuthMiddleware,
	tenantMiddleware *middleware.TenantMiddleware,
) {

	// ================= Handlers =================
	lessonHandlers := lessonbuilders.NewLessonHandlerBuilder(
		lessonUCs,
	)

	// ================= Routes =================
	protected := genericrouter.RegisterCRUDRoutes(
		r,
		"lessons",
		authMiddleware,
		tenantMiddleware,
		lessonHandlers.CRUDHandler,
	)

	protected.GET("/lessons/fetch", lessonHandlers.FetchLessonsHandler.GetLessons)
}
