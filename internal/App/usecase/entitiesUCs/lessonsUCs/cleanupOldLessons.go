package lessonsucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type CleanupOldLessonsUC struct {
	lessonQueryService entitiesrepos.LessonQueryService
}

type ICleanupOldLessonsUC interface {
	Execute(ctx context.Context) error
}

func NewCleanupOldLessonsUC(lessonQueryService entitiesrepos.LessonQueryService) ICleanupOldLessonsUC {
	return &CleanupOldLessonsUC{lessonQueryService}
}

func (uc *CleanupOldLessonsUC) Execute(ctx context.Context) error {
	return uc.lessonQueryService.CleanupOldLessons(ctx)
}
