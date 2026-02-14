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
