package lessonsucs

import (
	"context"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/lessonsUCs/lessonShedulesUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetLessonsUC struct {
	repo       entitiesrepos.LessonsRepo
	generateUC lessonshedulesucs.ICreateLessonsFromShceduleUC
	cleanupUC  ICleanupOldLessonsUC
}

type IGetLessonsUC interface {
	Execute(ctx context.Context) ([]entities.Lesson, error)
}

func NewGetLessonsUC(
	repo entitiesrepos.LessonsRepo,
	generateUC lessonshedulesucs.ICreateLessonsFromShceduleUC,
	cleanupUC ICleanupOldLessonsUC,
) IGetLessonsUC {
	return &GetLessonsUC{repo, generateUC, cleanupUC}
}

func (uc *GetLessonsUC) Execute(ctx context.Context) ([]entities.Lesson, error) {

	if err := uc.generateUC.Execute(ctx); err != nil {
		return nil, err
	}

	if err := uc.cleanupUC.Execute(ctx); err != nil {
		return nil, err
	}

	lessons := []entities.Lesson{}

	if err := uc.repo.GetAll(ctx, &lessons); err != nil {
		return nil, err
	}

	return lessons, nil
}
