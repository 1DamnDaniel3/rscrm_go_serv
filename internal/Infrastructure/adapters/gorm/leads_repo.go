package adapters

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
