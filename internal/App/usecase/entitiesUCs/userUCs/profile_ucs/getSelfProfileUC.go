package profileucs

import (
	"context"

	userprofilepolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/user_profile_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type GetSelfProfileUC struct {
	repo   entitiesrepos.UserProfileRepo
	policy userprofilepolicies.UserProfilePolicies
}

type IGetSelfProfileUC interface {
	Execute(ctx context.Context, id int64, profile *[]entities.UserProfile) error
}

func NewGetSelfProfileUC(
	repo entitiesrepos.UserProfileRepo,
	policy userprofilepolicies.UserProfilePolicies,
) IGetSelfProfileUC {
	return &GetSelfProfileUC{
		repo:   repo,
		policy: policy,
	}
}

func (uc *GetSelfProfileUC) Execute(ctx context.Context, acc_id int64, profile *[]entities.UserProfile) error {

	if err := uc.policy.ReadPolicies.ReadSelfProfile(ctx, acc_id); err != nil {
		return err
	}

	filter := make(map[string]interface{})
	filter["account_id"] = acc_id

	if err := uc.repo.GetAllWhere(ctx, filter, profile); err != nil {
		return err
	}

	return nil

}
