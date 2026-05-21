package useraccountpolicies

import (
	"context"
	"fmt"
	"slices"

	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type UserAccountReadPolicy struct{}

type IUserAccountReadPolicy interface {
	CanReadAllEmpByRole(ctx context.Context, roles []string) (*policytypes.Scope, error)
}

func NewUserAccountReadPolicy() IUserAccountReadPolicy {
	return &UserAccountReadPolicy{}
}

func (p *UserAccountReadPolicy) CanReadAllEmpByRole(
	ctx context.Context,
	roles []string,
) (*policytypes.Scope, error) {

	if slices.Contains(roles, "admin") {
		return nil, fmt.Errorf("Nah. You can't get admins. good luck ;)")
	}

	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// Admin → все аккаунты
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → все аккаунты своей школы
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// Manager только учителей
	if policyutils.HasAnyRole(user, valuetypes.Manager) {
		if slices.Contains(roles, "teacher") {
			return &policytypes.Scope{
				IsGlobal:  false,
				School_id: user.School_id,
			}, nil
		}
	}

	// Остальные получат только свой акк
	return &policytypes.Scope{
		IsGlobal: false,
		User_id:  user.ID,
	}, nil
}
