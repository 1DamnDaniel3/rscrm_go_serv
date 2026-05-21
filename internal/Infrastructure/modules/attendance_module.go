package modules

import (
	attendancepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/attendance_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type AttendanceModule struct {
	AttendanceRepo         entitiesrepos.AttendanceRepo
	AttendanceQueryService entitiesrepos.AttendanceQueryService
	AttendancePolicies     *attendancepolicies.AttendancePolicies
}

func NewAttendanceModule(
	db *gorm.DB,
) *AttendanceModule {
	return &AttendanceModule{
		AttendanceRepo: gormentityrepos.NewGormAttendanceRepo(db),

		AttendanceQueryService: gormentityrepos.NewGormAttendanceQueryService(db),

		AttendancePolicies: attendancepolicies.NewAttendancePolicies(
			attendancepolicies.NewAttendanceCrudPolicy(),
			attendancepolicies.NewAttendanceCreatePolicy(),
			attendancepolicies.NewAttendanceUpdatePolicy(),
		),
	}
}
