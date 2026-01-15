package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormLeadGroupRepo struct {
	*genericAdapter.GormRepository[entities.LeadGroup]
	db *gorm.DB
}

func NewGormLeadGroupsRepo(db *gorm.DB) entitiesrepos.LeadGroupsRepo {
	return &GormLeadGroupRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.LeadGroup](db),
		db:             db}
}
