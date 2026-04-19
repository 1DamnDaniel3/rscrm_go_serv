package modules

import (
	studentgrouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_group_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type StudentGroupsModule struct {
	StudentGroupsRepo    entitiesrepos.StudentGroupsRepo
	StudentGroupPolicies *studentgrouppolicies.StudentGroupPolicies
}

func NewStudentGroupsModule(
	db *gorm.DB,
) *StudentGroupsModule {
	return &StudentGroupsModule{
		StudentGroupsRepo: gormentityrepos.NewGormStudentGroupsRepo(db),

		StudentGroupPolicies: studentgrouppolicies.NewStudentGroupPolicies(
			studentgrouppolicies.NewStudentGroupCrudPolicy(),
		),
	}
}
