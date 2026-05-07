package entityroutes

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"

	schedulebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/schedule_builders"

	genericrouter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/routes/entity_routes/generic_router"

	"github.com/gin-gonic/gin"
)

func ScheduleRoutes(
	r *gin.RouterGroup,
	app *infrastructure.AppContainer,
	scheduleUCs *schedulebuilders.ScheduleUseCases,
) {

	// ================= Handlers =================
	scheduleHandlers := schedulebuilders.NewScheduleHandlerBuilder(
		scheduleUCs,
	)

	// ================= Routes =================
	genericrouter.RegisterCRUDRoutes(
		r,
		"schedules",
		app.AuthMiddleware,
		app.TenantMiddleware,
		scheduleHandlers.CRUDHandler,
	)
}
