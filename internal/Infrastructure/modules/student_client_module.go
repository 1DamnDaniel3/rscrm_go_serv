package modules

import (
	studentclientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/student_client_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type StudentClientModule struct {
	StudentClientsRepo    entitiesrepos.StudentClientsRepo
	StudentClientPolicies *studentclientpolicies.StudentClientPolicies
}

func NewStudentClientModule(
	db *gorm.DB,
) *StudentClientModule {
	return &StudentClientModule{
		StudentClientsRepo: gormentityrepos.NewGormStudentClientsRepo(db),

		StudentClientPolicies: studentclientpolicies.NewStudentClientPolicies(
			studentclientpolicies.NewStudentClientCrudPolicy(),
		),
	}
}
