package infrastructure

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/bcrypt"
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

	// -=== services
	Hasher ports.PasswordHasher
	JWT    ports.JWTservice
	Tx     services.Transaction
}

// -=== - === - === - === - constructor

func NewAppContainer(db *gorm.DB, hasher *bcrypt.BcryptHasher, jwt ports.JWTservice, tx services.Transaction) *AppContainer {
	return &AppContainer{
		// -=== modules
		SchoolModule:       modules.NewSchoolModule(db),
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

		// -=== services
		Hasher: hasher,
		JWT:    jwt,
		Tx:     tx,
	}
}
