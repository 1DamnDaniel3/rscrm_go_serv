package gormentityrepos

import (
	"context"
	"time"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	genericAdapter "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormutils"
	"gorm.io/gorm"
)

type GormLessonRepo struct {
	*genericAdapter.GormRepository[entities.Lesson]
	db *gorm.DB
}

func NewGormLessonRepo(db *gorm.DB) entitiesrepos.LessonsRepo {
	return &GormLessonRepo{
		GormRepository: genericAdapter.NewGormRepository[entities.Lesson](db),
		db:             db,
	}
}

// ---=========================================== QuerySerivce ===========================================---

type GormLessonQueryService struct {
	db *gorm.DB
}

func NewGormLessonQueryService(db *gorm.DB) entitiesrepos.LessonQueryService {
	return &GormLessonQueryService{db}
}

func (r *GormLessonQueryService) CleanupOldLessons(ctx context.Context) error {
	db := gormutils.DBFromCtx(ctx, r.db)
	db = gormutils.ApplyTenantFilter[entities.Lesson](ctx, db)

	twoAgo := time.Now().AddDate(0, -2, 0)

	return db.Where("lesson_date < ?", twoAgo).Delete(&entities.Lesson{}).Error
}
