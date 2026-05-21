package attendancebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
	studentbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_builders"

	attendancepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/attendance_policies"
	attendanceucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/attendanceUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type AttendanceUseCases struct {
	CRUD               genericcruduc.ICRUDUseCase[entities.Attendance]
	GenerateAttendance attendanceucs.IGenerateAttendanceUC
	MarkAttendance     attendanceucs.IMarkAttendanceUC
}

func NewAttendanceUseCases(
	tx services.Transaction,
	attendanceModule *modules.AttendanceModule,
	studentUCs *studentbuilders.StudentUseCases,
) *AttendanceUseCases {

	return &AttendanceUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			attendanceModule.AttendanceRepo,
			attendanceModule.AttendancePolicies.CRUD,
		),
		GenerateAttendance: attendanceucs.NewGenerateAttendanceUC(
			tx,
			attendanceModule.AttendanceRepo,
			attendanceModule.AttendanceQueryService,
			attendanceModule.AttendancePolicies.CreatePolicy,
			studentUCs.GroupedStudents,
		),
		MarkAttendance: attendanceucs.NewMarkAttendanceUC(
			attendanceModule.AttendanceRepo,
			attendancepolicies.NewAttendanceUpdatePolicy(),
		),
	}
}
