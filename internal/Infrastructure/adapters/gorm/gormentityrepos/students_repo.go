package gormentityrepos

import (
	"context"
	"fmt"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormStudentsRepo struct {
	*genericAdapter.GormRepository[entities.Student]
	db *gorm.DB
}

func NewGormStudentsRepo(db *gorm.DB) entitiesrepos.StudentsRepo {
	return &GormStudentsRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Student](db),
		db:             db,
	}
}

func (r *GormStudentsRepo) GetGroupedStudents(ctx context.Context, group_id int64, entities *[]entities.Student) error {
	db := r.DBFromCtx(ctx)

	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	return db.
		Table("students s").
		Select("s.*").
		Joins("JOIN student_groups sg ON sg.student_id = s.id").
		Joins("JOIN groups g ON g.id = sg.group_id").
		Where("s.school_id = ? AND g.id = ?", schoolID, group_id).
		Find(&entities).Error
}
