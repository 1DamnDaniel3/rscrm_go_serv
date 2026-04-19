package lessonbuilders

import (
	lessonsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs"
	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/lessonsUCs/lessonShedulesUCs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type LessonUseCases struct {
	CRUD                 genericcruduc.ICRUDUseCase[entities.Lesson]
	GenerateFromSchedule lessonshedulesucs.ICreateLessonsFromShceduleUC
	Cleanup              lessonsucs.ICleanupOldLessonsUC
	FetchLessons         lessonsucs.IGetLessonsUC
}

func NewLessonUseCasesBuilder(
	tx services.Transaction,
	lessonModule *modules.LessonModule,
	scheduleModule *modules.ScheduleModule,
) *LessonUseCases {

	return &LessonUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			lessonModule.LessonsRepo,
			lessonModule.LessonPolicies.CRUD,
		),

		// ================= GENERATE FROM SCHEDULE =================
		GenerateFromSchedule: lessonshedulesucs.NewCreateLessonsFromShceduleUC(
			lessonModule.LessonsRepo,
			scheduleModule.ScheduleRepo,
			lessonModule.LessonPolicies.CRUD,
			scheduleModule.SchedulePolicies.CRUD,
		),

		// ================= CLEANUP =================
		Cleanup: lessonsucs.NewCleanupOldLessonsUC(
			lessonModule.LessonQueryService,
		),

		// ================= FETCH LESSONS =================
		FetchLessons: lessonsucs.NewGetLessonsUC(
			lessonModule.LessonsRepo,
			lessonshedulesucs.NewCreateLessonsFromShceduleUC(
				lessonModule.LessonsRepo,
				scheduleModule.ScheduleRepo,
				lessonModule.LessonPolicies.CRUD,
				scheduleModule.SchedulePolicies.CRUD,
			),
			lessonsucs.NewCleanupOldLessonsUC(
				lessonModule.LessonQueryService,
			),
			lessonModule.LessonPolicies.CRUD,
		),
	}
}
