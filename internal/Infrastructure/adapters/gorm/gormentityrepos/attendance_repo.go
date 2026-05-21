package gormentityrepos

import (
	"context"
	"errors"
	"fmt"
	"time"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormAttendanceRepo struct {
	*generic.GormRepository[entities.Attendance]
	db *gorm.DB
}

func NewGormAttendanceRepo(db *gorm.DB) entitiesrepos.AttendanceRepo {
	return &GormAttendanceRepo{
		GormRepository: generic.NewGormRepository[entities.Attendance](db),
		db:             db,
	}
}

func (r *GormAttendanceRepo) MarkAttendance(
	ctx context.Context,
	attendance_id int64,
	scope *policytypes.Scope,
) (*entities.Attendance, error) {
	db := gormutils.DBFromCtx(ctx, r.db)

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	var updatedAttendance entities.Attendance

	err = db.Transaction(func(tx *gorm.DB) error {
		query := tx.Table("attendance AS a").
			Joins("JOIN lessons l ON l.id = a.lesson_id").
			Where("a.id = ?", attendance_id)

		if scope.User_id != 0 {
			query = query.Where("l.teacher_id = ?", scope.User_id)
		}

		if scope.School_id != "" {
			query = query.Where("a.school_id = ?", scope.School_id)
		}

		// сначала выбираем запись
		if err := query.First(&updatedAttendance).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("attendance not found or forbidden")
			}
			return err
		}

		// toggle status
		newStatus := "absent"
		switch updatedAttendance.Status {
		case "absent":
			newStatus = "presence"
		case "presence":
			newStatus = "absent"
		}

		// обновляем
		return tx.Model(&updatedAttendance).Updates(map[string]interface{}{
			"status":    newStatus,
			"marked_by": user.ID,
			"marked_at": time.Now(),
		}).Error
	})

	if err != nil {
		return nil, err
	}

	return &updatedAttendance, nil
}

// ========================= QueryService ==========================

type GormAttendanceQueryService struct {
	db *gorm.DB
}

func NewGormAttendanceQueryService(db *gorm.DB) entitiesrepos.AttendanceQueryService {
	return &GormAttendanceQueryService{
		db: db,
	}
}

func (s *GormAttendanceQueryService) FindAttendancies(
	ctx context.Context,
	student_ids []int64,
	lesson_id int64,
	scope *policytypes.Scope,
) ([]entities.Attendance, error) {

	db := gormutils.DBFromCtx(ctx, s.db)
	db, err := gormutils.ApplyScope(db, scope, "", "school_id")
	if err != nil {
		return nil, err
	}

	attendancies := []entities.Attendance{}
	err = db.Where("lesson_id = ? AND student_id IN ?", lesson_id, student_ids).
		Find(&attendancies).Error
	if err != nil {
		return nil, err
	}

	return attendancies, nil

}
