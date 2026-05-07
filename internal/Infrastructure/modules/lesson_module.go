package modules

import (
	lessonpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lesson_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type LessonModule struct {
	LessonsRepo        entitiesrepos.LessonsRepo
	LessonQueryService entitiesrepos.LessonQueryService
	LessonPolicies     *lessonpolicies.LessonPolicies
}

func NewLessonModule(
	db *gorm.DB,
) *LessonModule {
	return &LessonModule{
		LessonsRepo:        gormentityrepos.NewGormLessonRepo(db),
		LessonQueryService: gormentityrepos.NewGormLessonQueryService(db),
		LessonPolicies: lessonpolicies.NewLessonPolicies(
			lessonpolicies.NewLessonCrudPolicy(),
			lessonpolicies.NewLessonCreatePolicy(),
		),
	}
}
