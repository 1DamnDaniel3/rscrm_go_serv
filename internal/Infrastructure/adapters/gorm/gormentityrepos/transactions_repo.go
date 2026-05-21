package gormentityrepos

import (
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"gorm.io/gorm"
)

type GormTransactionRepo struct {
	db *gorm.DB
	*generic.GormRepository[entities.Transaction]
}

func NewGormTransactionRepo(db *gorm.DB) entitiesrepos.TransactionRepo {
	return &GormTransactionRepo{
		db:             db,
		GormRepository: generic.NewGormRepository[entities.Transaction](db),
	}
}
