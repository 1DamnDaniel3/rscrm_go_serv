package policyutils

import (
	"context"
	"fmt"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
)

func GetUserFromCtx(ctx context.Context) (*policytypes.User, error) {
	userCtx, ok := ctx.Value(contextkeys.User).(*valuetypes.UserContext)
	if !ok || userCtx == nil {
		return nil, fmt.Errorf("user context not found in context")
	}

	// Преобразуем роли в []valuetypes.Role
	roles := make([]valuetypes.Role, len(userCtx.Roles))
	for i, r := range userCtx.Roles {
		roles[i] = valuetypes.Role(r)
	}

	return &policytypes.User{
		ID:        userCtx.UserID,
		Roles:     roles,
		School_id: userCtx.SchoolID,
	}, nil
}
