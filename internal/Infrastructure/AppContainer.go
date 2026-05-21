package infrastructure

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/middleware"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
	"gorm.io/gorm"
)

type AppContainer struct {
	// -=== modules

	SchoolModule       *modules.SchoolModule
	AccountModule      *modules.AccountModule
	ProfileModule      *modules.ProfileModule
	AccountRolesModule *modules.AccountRolesModule
	RolesModule        *modules.RolesModule

	SourceModule *modules.SourceModule
	StatusModule *modules.StatusModule

	LeadModule          *modules.LeadModule
	ClientModule        *modules.ClientModule
	StudentClientModule *modules.StudentClientModule
	StudentModule       *modules.StudentModule

	GroupModule *modules.GroupModule

	LeadGroupsModule    *modules.LeadGroupsModule
	ClientGroupsModule  *modules.ClientGroupsModule
	StudentGroupsModule *modules.StudentGroupsModule

	ScheduleModule *modules.ScheduleModule
	LessonModule   *modules.LessonModule

	EmployeeRatePolicyModule *modules.EmployeeRatePolicyModule
	EmployeeRateRuleModule   *modules.EmployeeRateRuleModule
	EmployeeRateModule       *modules.EmployeeRateModule

	TransactionModule *modules.TransactionModule

	SalaryAccuralsModule     *modules.SalaryAccuralsModule
	SalaryAccuralItemsModule *modules.SalaryAccuralItemsModule

	SubscriptionModule         *modules.SubscriptionModule
	StudentSubscriptionsModule *modules.StudentSubscriptionsModule
	LessonSubscriptionsModule  *modules.LessonSubscriptionsModule

	AttendanceModule *modules.AttendanceModule

	DanceStyleModule *modules.DanceStyleModule

	// -=== services
	Hasher ports.PasswordHasher
	JWT    ports.JWTservice
	Tx     services.Transaction

	// -=== middleware
	AuthMiddleware   *middleware.AuthMiddleware
	TenantMiddleware *middleware.TenantMiddleware
}

// -=== - === - === - === - constructor

func NewAppContainer(
	db *gorm.DB,
	hasher ports.PasswordHasher,
	jwt ports.JWTservice,
	tx services.Transaction,
	AuthMiddleware *middleware.AuthMiddleware,
	TenantMiddleware *middleware.TenantMiddleware,
) *AppContainer {
	return &AppContainer{
		// -=== modules
		SchoolModule:       modules.NewSchoolModule(db),
		AccountModule:      modules.NewAccountModule(db, hasher),
		ProfileModule:      modules.NewProfileModule(db),
		AccountRolesModule: modules.NewAccountRolesModule(db),
		RolesModule:        modules.NewRolesModule(db),

		SourceModule: modules.NewSourceModule(db),
		StatusModule: modules.NewStatusModule(db),

		LeadModule:          modules.NewLeadModule(db),
		ClientModule:        modules.NewClientModule(db),
		StudentClientModule: modules.NewStudentClientModule(db),
		StudentModule:       modules.NewStudentModule(db),

		GroupModule: modules.NewGroupModule(db),

		LeadGroupsModule:    modules.NewLeadGroupsModule(db),
		ClientGroupsModule:  modules.NewClientGroupsModule(db),
		StudentGroupsModule: modules.NewStudentGroupsModule(db),

		ScheduleModule: modules.NewScheduleModule(db),
		LessonModule:   modules.NewLessonModule(db),

		EmployeeRatePolicyModule: modules.NewEmployeeRatePolicyModule(db),
		EmployeeRateRuleModule:   modules.NewEmployeeRateRuleModule(db),
		EmployeeRateModule:       modules.NewEmployeeRateModule(db),

		TransactionModule: modules.NewTransactionModule(db),

		SalaryAccuralsModule:     modules.NewSalaryAccuralsModule(db),
		SalaryAccuralItemsModule: modules.NewSalaryAccuralItemsModule(db),

		SubscriptionModule:         modules.NewSubscriptionModule(db),
		StudentSubscriptionsModule: modules.NewStudentSubscriptionsModule(db),
		LessonSubscriptionsModule:  modules.NewLessonSubscriptionsModule(db),

		AttendanceModule: modules.NewAttendanceModule(db),

		DanceStyleModule: modules.NewDanceStyleModule(db),

		// -=== services
		Hasher: hasher,
		JWT:    jwt,
		Tx:     tx,

		// -=== middleware

		AuthMiddleware:   AuthMiddleware,
		TenantMiddleware: TenantMiddleware,
	}
}
