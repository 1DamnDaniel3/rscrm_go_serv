package policyutils

import (
	"context"
	"fmt"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

func RequireRoles(ctx context.Context, roles ...valuetypes.Role) error {
	user, err := GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if !HasAnyRole(user, roles...) {
		return fmt.Errorf("forbidden. haven't required roles in ctx")
	}

	return nil
}
