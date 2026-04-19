package builders

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	clientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/client_builders"
	groupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/group_builders"
	leadbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_builders"
	leadgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_group_builders"
	lessonbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lesson_builders"
	schedulebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/schedule_builders"
	schoolbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/school_builders"
	sourcebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/source_builders"
	statusbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/status_builders"
	studentbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_builders"
	studentclientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_clients_builders"
	studentgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_group_builders"
	useraccountbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_accounts_builders"
	userprofilebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_profile_builders"
)

type UseCases struct {
	Client         *clientbuilders.ClientUseCases
	Group          *groupbuilders.GroupUseCases
	Lead           *leadbuilders.LeadUseCases
	LeadGroup      *leadgroupbuilders.LeadGroupsUseCaseBuilder
	Lesson         *lessonbuilders.LessonUseCases
	Schedule       *schedulebuilders.ScheduleUseCases
	School         *schoolbuilders.SchoolUseCases
	Source         *sourcebuilders.SourceUseCases
	Status         *statusbuilders.StatusUseCases
	Student        *studentbuilders.StudentUseCases
	StudentGroup   *studentgroupbuilders.StudentGroupUseCases
	StudentClients *studentclientbuilders.StudentClientUseCases
	UserAccount    *useraccountbuilders.UserUseCases
	UserProfile    *userprofilebuilders.UserProfileUseCases
}

func NewUseCases(app *infrastructure.AppContainer) *UseCases {
	return &UseCases{
		Client: clientbuilders.NewClientUseCases(
			app.Tx,
			app.ClientModule,
			app.ClientGroupsModule,
		),

		Group: groupbuilders.NewGroupUseCases(
			app.GroupModule.GroupRepo,
			app.GroupModule.GroupPolicies,
		),

		Lead: leadbuilders.NewLeadUseCasesBuilder(
			app.Tx,
			app.LeadModule,
			app.LeadGroupsModule,
		),

		LeadGroup: leadgroupbuilders.NewLeadGroupsUseCaseBuilder(
			app.LeadGroupsModule,
		),

		Lesson: lessonbuilders.NewLessonUseCasesBuilder(
			app.Tx,
			app.LessonModule,
			app.ScheduleModule,
		),

		Schedule: schedulebuilders.NewScheduleUseCasesBuilder(
			app.Tx,
			app.ScheduleModule,
		),

		School: schoolbuilders.NewSchoolUseCasesBuilder(
			app.SchoolModule,
		),

		Source: sourcebuilders.NewSourceUseCasesBuilder(
			app.SourceModule,
		),

		Status: statusbuilders.NewStatusUseCasesBuilder(
			app.StatusModule,
		),

		Student: studentbuilders.NewStudentUseCasesBuilder(
			app.Tx,
			app.StudentModule,
			app.StudentGroupsModule,
		),

		StudentClients: studentclientbuilders.NewStudentClientUseCasesBuilder(
			app.Tx,
			app.StudentClientModule,
			app.ClientModule,
		),

		UserAccount: useraccountbuilders.NewUserUseCasesBuilder(
			app.Tx,
			app.Hasher,
			app.JWT,
			app.AccountModule,
			app.ProfileModule,
			app.AccountRolesModule,
			app.RolesModule,
			app.SchoolModule,
		),

		UserProfile: userprofilebuilders.NewUserProfileUseCasesBuilder(
			app.ProfileModule,
		),
	}
}
