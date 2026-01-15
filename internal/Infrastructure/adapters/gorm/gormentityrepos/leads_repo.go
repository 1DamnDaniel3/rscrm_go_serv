package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
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
func (r *GormLeadsRepo) GetGroupedLeads(school_id string, group_id int64, entities *[]entities.Lead) error {
	return r.db.
		Table("leads l").
		Select("l.*").
		Joins("JOIN lead_groups lg ON lg.lead_id = l.id").
		Joins("JOIN groups g ON g.id = lg.group_id").
		Where("l.school_id = ? AND g.id = ?", school_id, group_id).
		Find(&entities).Error
}
