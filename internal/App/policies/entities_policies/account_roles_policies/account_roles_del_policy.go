package accountrolespolicies

import (
	"context"
	"fmt"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type AccountRolesDeletePolicy struct{}

type IAccountRolesDeletePolicy interface {
	CanRemoveRole(ctx context.Context, role_id, acc_id int64, target_acc_school_id string) (*policytypes.Scope, error)
}

func NewAccountRolesDeletePolicy() IAccountRolesDeletePolicy {
	return &AccountRolesDeletePolicy{}
}

func (p *AccountRolesDeletePolicy) CanRemoveRole(
	ctx context.Context,
	role_id,
	acc_id int64,
	target_acc_school_id string,
) (*policytypes.Scope, error) {

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		if acc_id == user.ID {
			return nil, fmt.Errorf("you cannot change your roles.")
		}
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		if role_id == 1 || role_id == 2 {
			return nil, fmt.Errorf("you cannot remove this role. Access denied.")
		}

		if target_acc_school_id != user.School_id {
			return nil, fmt.Errorf("Access denied. Check account_id")
		}

		if acc_id == user.ID {
			return nil, fmt.Errorf("you cannot remove your roles.")
		}

		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("access denied")

}
