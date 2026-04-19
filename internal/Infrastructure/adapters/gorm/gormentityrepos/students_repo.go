package gormentityrepos

import (
	"context"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
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

	school_id, err := gormutils.GetTenandID(ctx)
	if err != nil {
		return err
	}

	return db.
		Table("students s").
		Select("s.*").
		Joins("JOIN student_groups sg ON sg.student_id = s.id").
		Joins("JOIN groups g ON g.id = sg.group_id").
		Where("s.school_id = ? AND g.id = ?", school_id, group_id).
		Find(&entities).Error
}

// ========================= QUERY SERVICE

type GormStudentQueryService struct {
	db *gorm.DB
}

func NewGormStudentQueryService(db *gorm.DB) entitiesrepos.StudentQueryService {
	return &GormStudentQueryService{db}
}

func (r *GormStudentQueryService) Search(ctx context.Context, input string, studentSlice *[]entities.Student) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	school_id, err := gormutils.GetTenandID(ctx)
	if err != nil {
		return err
	}
	return db.
		Table("students s").
		Where("s.school_id = ? AND (s.name ILIKE ?)", school_id, "%"+input+"%").
		Find(studentSlice).Error
}

func (r *GormStudentQueryService) GetStudentClients(ctx context.Context, student_id int64, clientsSlice *[]businessobjects.GetStudentClientsBO) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	school_id, err := gormutils.GetTenandID(ctx)
	if err != nil {
		return err
	}

	return db.
		Table("students s").
		Select("c.*, sc.is_payer, sc.relation, sc.id AS relation_id").
		Joins("JOIN student_clients sc ON s.id = sc.student_id").
		Joins("JOIN clients c ON sc.client_id = c.id").
		Where("s.id = ? AND sc.school_id = ?", student_id, school_id).
		Scan(clientsSlice).Error
}

// get student groups
func (r *GormStudentQueryService) GetStudentGroups(ctx context.Context,
	student_id int64, groupSlice *[]entities.Group) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	// db = gormutils.ApplyTenantFilter[entities.Client](ctx, db)
	school_id, err := gormutils.GetTenandID(ctx)
	if err != nil {
		return err
	}
	return db.
		Table("student_groups sg").
		Select("g.*").
		Joins("JOIN groups g ON g.id = sg.group_id").
		Where("sg.student_id = ? AND g.school_id = ?", student_id, school_id).
		Find(groupSlice).Error
}
