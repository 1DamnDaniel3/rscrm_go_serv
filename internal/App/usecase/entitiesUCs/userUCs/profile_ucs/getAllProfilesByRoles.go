package profileucs

import (
	"context"

	useraccountpolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_account_policies.go"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetAllProfilesByRolesUC struct {
	profileRepo entitiesrepos.UserProfileRepo
	policy      useraccountpolicies.IUserAccountReadPolicy
}

type IGetAllProfilesByRolesUC interface {
	Execute(ctx context.Context, roles []string) ([]entities.UserProfile, error)
}

func NewGetAllProfilesByRolesUC(
	profileRepo entitiesrepos.UserProfileRepo,
	policy useraccountpolicies.IUserAccountReadPolicy,
) IGetAllProfilesByRolesUC {
	return &GetAllProfilesByRolesUC{profileRepo, policy}
}

func (uc *GetAllProfilesByRolesUC) Execute(ctx context.Context, roles []string) ([]entities.UserProfile, error) {

	scope, err := uc.policy.CanReadAllEmpByRole(ctx, roles)
	if err != nil {
		return nil, err
	}

	profiles, err := uc.profileRepo.GetAllProfilesByRoles(ctx, scope, roles...)
	if err != nil {
		return nil, err
	}

	return *profiles, nil

}
