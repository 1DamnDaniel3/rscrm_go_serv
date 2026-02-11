package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormStudentGroupsRepo struct {
	*genericAdapter.GormRepository[entities.StudentGroup]
	db *gorm.DB
}

func NewGormStudentGroupsRepo(db *gorm.DB) entitiesrepos.StudentGroupsRepo {
	return &GormStudentGroupsRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.StudentGroup](db),
		db:             db,
	}
}
