package userprofilepolicies

import (
	"context"
	"fmt"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
)

type ReadProfilePolicy struct{}

type IReadProfilePolicy interface {
	ReadSelfProfile(ctx context.Context, requiredID int64) error
}

func NewReadProfilePolicy() IReadProfilePolicy {
	return &ReadProfilePolicy{}
}

// ---====== Methods ======---

func (p *ReadProfilePolicy) ReadSelfProfile(ctx context.Context, requiredID int64) error {

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if user.ID != requiredID {
		return fmt.Errorf("you can read only your own profile data. err from policy layer")
	}

	return nil
}
