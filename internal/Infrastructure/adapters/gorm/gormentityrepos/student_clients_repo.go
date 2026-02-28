package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormStudentClientsRepo struct {
	*genericAdapter.GormRepository[entities.StudentClient]
	db *gorm.DB
}

func NewGormStudentClientsRepo(db *gorm.DB) entitiesrepos.StudentClientsRepo {
	return &GormStudentClientsRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.StudentClient](db),
		db:             db,
	}
}
