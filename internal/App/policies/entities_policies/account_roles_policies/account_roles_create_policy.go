package accountrolespolicies

import (
	"context"
	"fmt"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type AccountRolesCreatePolicy struct{}

type IAccountRolesCreatePolicy interface {
	CanAssignRole(ctx context.Context, role_id int64, acc_id int64, target_acc_school_id string) (*policytypes.Scope, error)
}

func NewAccountRolesCreatePolicy() IAccountRolesCreatePolicy {
	return &AccountRolesCreatePolicy{}
}

func (p *AccountRolesCreatePolicy) CanAssignRole(
	ctx context.Context,
	role_id int64,
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
			return nil, fmt.Errorf("you cannot assign this role. Access denied.")
		}

		if target_acc_school_id != user.School_id {
			return nil, fmt.Errorf("Access denied. Check account_id")
		}

		if acc_id == user.ID {
			return nil, fmt.Errorf("you cannot change your roles.")
		}

		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	return nil, fmt.Errorf("access denied")
}
