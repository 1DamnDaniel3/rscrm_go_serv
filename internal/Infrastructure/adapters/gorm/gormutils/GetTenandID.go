package gormutils

import (
	"context"
	"fmt"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
)

// Just get school_id from request ctx
func GetTenandID(ctx context.Context) (string, error) {
	userCtx, ok := ctx.Value(contextkeys.User).(*valuetypes.UserContext)
	if !ok || userCtx.SchoolID == "" {
		return "", fmt.Errorf("School Id is missing. gormutils.GetTenandID")

	}

	return userCtx.SchoolID, nil
}
