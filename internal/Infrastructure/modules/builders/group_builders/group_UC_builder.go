package groupbuilders

import (
	grouppolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/group_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GroupUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.Group]
}

func NewGroupUseCases(
	groupRepo entitiesrepos.GroupRepo,
	groupPolicies *grouppolicies.GroupPolicies,
) *GroupUseCases {

	return &GroupUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			groupRepo,
			groupPolicies.CRUD,
		),
	}
}
