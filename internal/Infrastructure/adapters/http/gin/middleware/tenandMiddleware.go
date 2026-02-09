package middleware

import (
	"context"
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/contextkeys"
	"github.com/gin-gonic/gin"
)

type TenantMiddleware struct{}

func NewTenandMiddleware() *TenantMiddleware {
	return &TenantMiddleware{}
}

func (t *TenantMiddleware) TryTenand() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get school_id from token
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

		schoolID, ok := claims["school_id"].(string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "school_id not found in claims"})
			return
		}

		ctx := context.WithValue(
			c.Request.Context(),
			contextkeys.SchoolID,
			schoolID,
		)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
