package entitiesrepos

import (
	"context"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type LessonsRepo interface {
	genericrepo.Repository[entities.Lesson]
}

type LessonQueryService interface {
	CleanupOldLessons(ctx context.Context) error
}
