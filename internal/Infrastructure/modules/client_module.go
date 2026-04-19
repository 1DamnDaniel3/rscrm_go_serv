package modules

import (
	clientpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/client_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/gorm/gormentityrepos"
	"gorm.io/gorm"
)

type ClientModule struct {
	ClientRepo         entitiesrepos.ClientRepo
	ClientQueryService entitiesrepos.ClientsQueryService
	ClientPolicies     *clientpolicies.ClientPolicies
}

func NewClientModule(
	db *gorm.DB,
) *ClientModule {
	return &ClientModule{
		ClientRepo: gormentityrepos.NewGormClientRepo(db),

		ClientQueryService: gormentityrepos.NewGormClientQueryService(db),

		ClientPolicies: clientpolicies.NewClientPolicies(
			clientpolicies.NewClientCrudPolicy(),
		),
	}
}
