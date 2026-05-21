package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"
	attendancebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/attendance_builders"
	"github.com/gin-gonic/gin"
)

func AttendanceRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	attendanceUCs *attendancebuilders.AttendanceUseCases,
) {

	// ================= Handlers =================
	attendanceHandlers := attendancebuilders.NewAttendanceHandlerBuilder(attendanceUCs, app.ClientModule)

	// ================= Routes =================
	protected := genericrouter.GetProtectedRouterGroup(
		r,
		app.AuthMiddleware,
		app.TenantMiddleware,
	)

	// protected.GET("/attendances", attendanceHandlers.CRUDHandler.GetAll)
	protected.POST("/attendances/generate", attendanceHandlers.GenerateHandler.Generate)
	protected.PATCH("/attendances/mark/:attendanceID", attendanceHandlers.MarkHandler.Mark)
}
