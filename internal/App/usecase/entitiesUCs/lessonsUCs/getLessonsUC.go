package lessonsucs

import (
	"context"
	"fmt"

	lessonpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/lesson_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs/lessonShedulesUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetLessonsUC struct {
	repo       entitiesrepos.LessonsRepo
	generateUC lessonshedulesucs.ICreateLessonsFromShceduleUC
	cleanupUC  ICleanupOldLessonsUC

	policy lessonpolicies.ILessonCrudPolicy
}

type IGetLessonsUC interface {
	Execute(ctx context.Context) ([]entities.Lesson, error)
}

func NewGetLessonsUC(
	repo entitiesrepos.LessonsRepo,
	generateUC lessonshedulesucs.ICreateLessonsFromShceduleUC,
	cleanupUC ICleanupOldLessonsUC,
	policy lessonpolicies.ILessonCrudPolicy,
) IGetLessonsUC {
	return &GetLessonsUC{repo, generateUC, cleanupUC, policy}
}

func (uc *GetLessonsUC) Execute(ctx context.Context) ([]entities.Lesson, error) {

	scope, err := uc.policy.CanReadAll(ctx)
	if err != nil {
		return nil, err
	}

	if err := uc.generateUC.Execute(ctx); err != nil {
		return nil, fmt.Errorf("lesson generating error: %v", err)
	}

	// ПОЗЖЕ НАСТРОИТЬ АРХИВАЦИЮ СТАРЫХ ЗАПИСЕЙ ВМЕСТО СТИРАНИЯ
	// if err := uc.cleanupUC.Execute(ctx); err != nil {
	// 	return nil, fmt.Errorf("cleanup error: %v", err)
	// }

	lessons := []entities.Lesson{}

	if err := uc.repo.GetAll(ctx, &lessons, scope); err != nil {
		return nil, fmt.Errorf("read err: %v", err)
	}

	return lessons, nil
}
