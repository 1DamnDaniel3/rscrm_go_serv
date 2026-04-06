package middleware

import (
	"context"
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	"github.com/gin-gonic/gin"
)

type TenantMiddleware struct{}

func NewTenandMiddleware() *TenantMiddleware {
	return &TenantMiddleware{}
}

func (t *TenantMiddleware) TryTenand() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get claims from c.Context
		user, exists := c.Get("user")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
			return
		}

		claims, ok := user.(map[string]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data in context"})
			return
		}

		// == GET CLAIMS
		userID, ok := claims["id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "id not found in claims"})
			return
		}

		schoolID, ok := claims["school_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "school_id not found in claims"})
			return
		}

		rolesInf, ok := claims["roles"].([]interface{})
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "roles not found in claims"})
			return
		}

		roles := make([]string, len(rolesInf))
		for i, r := range rolesInf {
			str, ok := r.(string)
			if !ok {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid role type"})
				return
			}
			roles[i] = str
		}

		// build context

		ctx := context.WithValue(
			c.Request.Context(),
			contextkeys.User,
			&valuetypes.UserContext{
				UserID:   int64(userID),
				SchoolID: schoolID,
				Roles:    roles,
			},
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
