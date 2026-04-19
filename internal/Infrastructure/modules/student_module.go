package modules

import (
	studentpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type StudentModule struct {
	StudentRepo         entitiesrepos.StudentsRepo
	StudentQueryService entitiesrepos.StudentQueryService
	StudentPolicies     *studentpolicies.StudentPolicies
}

func NewStudentModule(
	db *gorm.DB,
) *StudentModule {
	return &StudentModule{
		StudentRepo: gormentityrepos.NewGormStudentsRepo(db),

		StudentQueryService: gormentityrepos.NewGormStudentQueryService(db),

		StudentPolicies: studentpolicies.NewStudentPolicies(
			studentpolicies.NewStudentCrudPolicy(),
		),
	}
}
