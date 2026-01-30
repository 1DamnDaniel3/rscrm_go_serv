package middleware

import (
	"net/http"
	"strings"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	JwtService ports.JWTservice
}

func NewAuthMiddleware(JwtService ports.JWTservice) *AuthMiddleware {
	return &AuthMiddleware{JwtService}
}

func (a *AuthMiddleware) TryAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("jwt")
		if err != nil {
			authHeader := c.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
			return
		}

		claims, err := a.JwtService.Verify(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("user", claims)
		c.Next()
	}
}
