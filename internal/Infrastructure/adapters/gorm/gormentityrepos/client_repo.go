package gormentityrepos

import (
	"context"
	"fmt"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormClientRepo struct {
	*genericAdapter.GormRepository[entities.Client]
	db *gorm.DB
}

func NewGormClientRepo(db *gorm.DB) entitiesrepos.ClientRepo {
	return &GormClientRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Client](db),
		db:             db,
	}
}

func (r *GormClientRepo) GetGroupedClients(ctx context.Context, group_id int64, entities *[]entities.Client) error {
	db := r.DBFromCtx(ctx)

	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	return db.
		Table("clients c").
		Select("c.*").
		Joins("JOIN client_groups cg ON cg.client_id = c.id").
		Joins("JOIN groups g ON g.id = cg.group_id").
		Where("c.school_id = ? AND g.id = ?", schoolID, group_id).
		Find(&entities).Error
}

// ========================= QUERY SERVICE

type GormClientQueryService struct {
	db *gorm.DB
}

func NewGormClientQueryService(db *gorm.DB) entitiesrepos.ClientsQueryService {
	return &GormClientQueryService{db}
}

func (r *GormClientQueryService) Search(ctx context.Context, input string, clientSlice *[]entities.Client) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}
	return db.
		Table("clients c").
		Where("c.school_id = ? AND (c.name ILIKE ?)", schoolID, "%"+input+"%").
		Find(clientSlice).Error
}

func (r *GormClientQueryService) GetClientGroups(ctx context.Context,
	client_id int64, groupSlice *[]entities.Group) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	// db = gormutils.ApplyTenantFilter[entities.Client](ctx, db)
	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}
	return db.
		Table("client_groups cg").
		Select("g.*").
		Joins("JOIN groups g ON g.id = cg.group_id").
		Where("cg.client_id = ? AND g.school_id = ?", client_id, schoolID).
		Find(groupSlice).Error
}

func (r *GormClientQueryService) GetClientStudents(
	ctx context.Context,
	client_id int64,
	studentsSlice *[]businessobjects.GetClientStudentsBO) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	return db.
		Table("clients c").
		Select("s.*, sc.id AS relation_id, sc.relation").
		Joins("JOIN student_clients sc ON c.id = sc.client_id").
		Joins("JOIN students s ON sc.student_id = s.id").
		Where("c.id = ? AND sc.school_id = ?", client_id, schoolID).
		Scan(studentsSlice).Error
}
