package policyutils

import (
	"context"
	"fmt"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type User struct {
	ID    int64
	Roles []valuetypes.Role
}

func GetUserFromCtx(ctx context.Context) (*User, error) {
	id, ok := ctx.Value("id").(int64)
	if !ok {
		return nil, fmt.Errorf("user_id not found in context. policy_utils layer error")
	}
	roles, ok := ctx.Value("roles").([]valuetypes.Role)
	if !ok {
		return nil, fmt.Errorf("user_roles not found in context. policy_utils layer error")
	}

	return &User{ID: id, Roles: roles}, nil
}
