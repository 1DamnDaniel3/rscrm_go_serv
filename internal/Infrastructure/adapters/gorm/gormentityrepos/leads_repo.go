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

type GormLeadsRepo struct {
	*genericAdapter.GormRepository[entities.Lead]
	db *gorm.DB
}

func NewGormLeadsRepo(db *gorm.DB) entitiesrepos.LeadsRepository {
	return &GormLeadsRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Lead](db),
		db:             db}
}

// Getting all leads of concrete school and using group filter
func (r *GormLeadsRepo) GetGroupedLeads(ctx context.Context, group_id int64, entities *[]entities.Lead) error {

	db := r.DBFromCtx(ctx)

	schoolID, ok := ctx.Value(contextkeys.SchoolID).(string)
	if !ok {
		return fmt.Errorf("school_id not found in context")
	}

	return db.
		Table("leads l").
		Select("l.*").
		Joins("JOIN lead_groups lg ON lg.lead_id = l.id").
		Joins("JOIN groups g ON g.id = lg.group_id").
		Where("l.school_id = ? AND g.id = ?", schoolID, group_id).
		Find(&entities).Error
}
