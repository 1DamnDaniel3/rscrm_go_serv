package attendancebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	attendancehandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/attendance_handlers"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type AttendanceHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Attendance,
		dto.AttendanceCreateUpdateDTO,
		dto.AttendanceResponseDTO,
	]
	GenerateHandler *attendancehandlers.GenerateAttendanciesHandler
	MarkHandler     *attendancehandlers.MarkAttendanceHandler
}

func NewAttendanceHandlerBuilder(useCases *AttendanceUseCases, modules *modules.ClientModule) *AttendanceHandlerBuilder {
	return &AttendanceHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Attendance,
			dto.AttendanceCreateUpdateDTO,
			dto.AttendanceResponseDTO,
		](useCases.CRUD),
		// ==================== GENERATE ==============
		GenerateHandler: attendancehandlers.NewGenerateAttendanciesHandler(
			useCases.GenerateAttendance,
		),

		// ==================== MARK ==============
		MarkHandler: attendancehandlers.NewMarkAttendanceHandler(
			useCases.MarkAttendance,
		),
	}
}
