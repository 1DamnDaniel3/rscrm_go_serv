package entitiesrepos

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type AttendanceRepo interface {
	genericrepo.Repository[entities.Attendance]
	MarkAttendance(ctx context.Context, attendance_id int64, scope *policytypes.Scope) (*entities.Attendance, error)
}

type AttendanceQueryService interface {
	// CleanupOldAttendancies(ctx context.Context) error
	FindAttendancies(ctx context.Context, student_ids []int64,
		lesson_id int64, scope *policytypes.Scope,
	) ([]entities.Attendance, error)
}
