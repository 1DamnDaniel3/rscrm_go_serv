package builders

import (
	infrastructure "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure"
	accountrolesbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/account_roles_builders"
	attendancebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/attendance_builders"
	clientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/client_builders"
	dancestylebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/dance_style_builders"
	groupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/group_builders"
	leadbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_builders"
	leadgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lead_group_builders"
	lessonbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lesson_builders"
	lessonsubscriptionsbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/lesson_subscription_builders"
	salaryaccuralitemsbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/salary_accural_items_builders"
	salaryaccuralsbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/salary_accurals_builders"
	schedulebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/schedule_builders"
	schoolbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/school_builders"
	sourcebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/source_builders"
	statusbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/status_builders"
	studentbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_builders"
	studentclientbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_clients_builders"
	studentgroupbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_group_builders"
	studentsubscriptionbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/student_subscriptions_builders"
	subscriptionbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/subscription_builders"
	transactionsbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/transactions_builders"
	useraccountbuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_accounts_builders"
	userprofilebuilders "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules/builders/user_profile_builders"
)

type UseCases struct {
	AccountRoles *accountrolesbuilders.AccountRolesUseCases
	Client       *clientbuilders.ClientUseCases
	Group        *groupbuilders.GroupUseCases
	Lead         *leadbuilders.LeadUseCases
	LeadGroup    *leadgroupbuilders.LeadGroupsUseCaseBuilder
	Lesson       *lessonbuilders.LessonUseCases
	Schedule     *schedulebuilders.ScheduleUseCases
	School       *schoolbuilders.SchoolUseCases
	Source       *sourcebuilders.SourceUseCases
	Status       *statusbuilders.StatusUseCases
	Student      *studentbuilders.StudentUseCases

	Subscription         *subscriptionbuilders.SubscriptionUseCases
	StudentSubscriptions *studentsubscriptionbuilders.StudentSubscriptionsUseCases
	LessonSubscriptions  *lessonsubscriptionsbuilders.LessonSubscriptionsUseCases

	Transactions *transactionsbuilders.TransactionsUseCases

	SalaryAccuralsUseCases     *salaryaccuralsbuilders.SalaryAccuralsUseCases
	SalaryAccuralItemsUseCases *salaryaccuralitemsbuilders.SalaryAccuralItemsUseCases

	StudentGroup   *studentgroupbuilders.StudentGroupUseCases
	StudentClients *studentclientbuilders.StudentClientUseCases
	UserAccount    *useraccountbuilders.UserUseCases
	UserProfile    *userprofilebuilders.UserProfileUseCases
	Attendance     *attendancebuilders.AttendanceUseCases
	DanceStyle     *dancestylebuilders.DanceStyleUseCases
}

func NewUseCases(app *infrastructure.AppContainer) *UseCases {

	// ================= ACCOUNT ROLES =================
	AccountRoles := accountrolesbuilders.NewAccountRolesUseCases(
		app.AccountRolesModule.AccountRolesRepo,
		app.AccountRolesModule.AccountRolePolicies,

		app.AccountModule.UserRepo,
		app.AccountModule.AccountPolicies,
	)

	// ================= CLIENT =================
	Client := clientbuilders.NewClientUseCases(
		app.Tx,
		app.ClientModule,
		app.ClientGroupsModule,
	)

	// ================= GROUP =================
	Group := groupbuilders.NewGroupUseCases(
		app.GroupModule.GroupRepo,
		app.GroupModule.GroupPolicies,
	)

	// ================= LEAD =================
	Lead := leadbuilders.NewLeadUseCasesBuilder(
		app.Tx,
		app.LeadModule,
		app.LeadGroupsModule,
	)

	// ================= LEAD GROUP =================
	LeadGroup := leadgroupbuilders.NewLeadGroupsUseCaseBuilder(
		app.Tx,
		app.LeadModule,
		app.LeadGroupsModule,
	)

	// ================= LESSON =================
	Lesson := lessonbuilders.NewLessonUseCasesBuilder(
		app.Tx,
		app.LessonModule,
		app.ScheduleModule,
	)

	// ================= SCHEDULE =================
	Schedule := schedulebuilders.NewScheduleUseCasesBuilder(
		app.Tx,
		app.ScheduleModule,
	)

	// ================= SCHOOL =================
	School := schoolbuilders.NewSchoolUseCasesBuilder(
		app.SchoolModule,
	)

	// ================= SOURCE =================
	Source := sourcebuilders.NewSourceUseCasesBuilder(
		app.SourceModule,
	)

	// ================= STATUS =================
	Status := statusbuilders.NewStatusUseCasesBuilder(
		app.StatusModule,
	)

	// ================= STUDENT =================
	Student := studentbuilders.NewStudentUseCasesBuilder(
		app.Tx,
		app.StudentModule,
		app.StudentGroupsModule,
	)

	// ================= SUBSCRIPTION =================
	Subscription := subscriptionbuilders.NewSubscriptionUCBuilder(
		app.SubscriptionModule,
	)

	// ================= STUDENT_SUBSCRIPTION =================
	StudentSubscriptions := studentsubscriptionbuilders.NewStudentClientUseCasesBuilder(
		app.Tx,
		app.StudentSubscriptionsModule,
	)
	// ================= LESSON_SUBSCRIPTION =================
	LessonSubscriptions := lessonsubscriptionsbuilders.NewLessonSubscriptionsUseCases(
		app.Tx,
		app.LessonSubscriptionsModule,
	)

	// ================= TRANSACTIONS =================
	Transactions := transactionsbuilders.NewSubscriptionUCBuilder(
		app.TransactionModule,
	)

	// ================= SALARY_ACCURALS =================
	SalaryAccuralsUseCases := salaryaccuralsbuilders.NewLessonSubscriptionsUseCases(
		app.Tx,
		app.SalaryAccuralsModule,
	)

	SalaryAccuralItemsUseCases := salaryaccuralitemsbuilders.NewLessonSubscriptionsUseCases(
		app.Tx,
		app.SalaryAccuralItemsModule,
	)

	// ================= STUDENT GROUP =================
	StudentGroup := studentgroupbuilders.NewStudentGroupUseCasesBuilder(
		app.StudentGroupsModule,
	)

	// ================= STUDENT CLIENTS =================
	StudentClients := studentclientbuilders.NewStudentClientUseCasesBuilder(
		app.Tx,
		app.StudentClientModule,
		app.ClientModule,
	)

	// ================= USER ACCOUNT =================
	UserAccount := useraccountbuilders.NewUserUseCasesBuilder(
		app.Tx,
		app.Hasher,
		app.JWT,
		app.AccountModule,
		app.ProfileModule,
		app.AccountRolesModule,
		app.RolesModule,
		app.SchoolModule,
	)

	// ================= USER PROFILE =================
	UserProfile := userprofilebuilders.NewUserProfileUseCasesBuilder(
		app.ProfileModule,
		app.AccountModule,
	)

	// ================= ATTENDANCE =================
	Attendance := attendancebuilders.NewAttendanceUseCases(
		app.Tx,
		app.AttendanceModule,
		Student,
	)

	// ================= DANCE STYLE =================

	DanceStyle := dancestylebuilders.NewDanceStyleUseCases(
		app.Tx,
		app.DanceStyleModule,
	)

	return &UseCases{
		AccountRoles: AccountRoles,
		Client:       Client,
		Group:        Group,
		Lead:         Lead,
		LeadGroup:    LeadGroup,
		Lesson:       Lesson,
		Schedule:     Schedule,
		School:       School,
		Source:       Source,
		Status:       Status,
		Student:      Student,

		Subscription:         Subscription,
		StudentSubscriptions: StudentSubscriptions,
		LessonSubscriptions:  LessonSubscriptions,

		Transactions: Transactions,

		SalaryAccuralsUseCases:     SalaryAccuralsUseCases,
		SalaryAccuralItemsUseCases: SalaryAccuralItemsUseCases,

		StudentGroup:   StudentGroup,
		StudentClients: StudentClients,
		UserAccount:    UserAccount,
		UserProfile:    UserProfile,
		Attendance:     Attendance,
		DanceStyle:     DanceStyle,
	}
}
