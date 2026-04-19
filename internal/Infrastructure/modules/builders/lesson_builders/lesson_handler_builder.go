package lessonbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	lessonhandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/lessonHandlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type LessonHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.Lesson,
		dto.LessonCreateUpdateDTO,
		dto.LessonResponseDTO,
	]

	FetchLessonsHandler *lessonhandlers.GetLessonsHandler
}

func NewLessonHandlerBuilder(
	uc *LessonUseCases,
) *LessonHandlerBuilder {

	return &LessonHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.Lesson,
			dto.LessonCreateUpdateDTO,
			dto.LessonResponseDTO,
		](uc.CRUD),

		// ================= FETCH =================
		FetchLessonsHandler: lessonhandlers.NewGetLessonsHandler(
			uc.FetchLessons,
		),
	}
}
