package policyutils

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

func HasAnyRole(user *policytypes.User, allowed ...valuetypes.Role) bool {
	for _, r := range user.Roles {
		for _, a := range allowed {
			if r == a {
				return true
			}
		}
	}
	return false
}
