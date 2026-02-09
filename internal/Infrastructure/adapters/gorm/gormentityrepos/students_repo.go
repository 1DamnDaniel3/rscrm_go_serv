package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
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
