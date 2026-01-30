package user

import (
	"net/http"
	"strings"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/gin-gonic/gin"
)

type AuthCheckHandler struct {
	JwtService ports.JWTservice
}

func NewAuthCheckHandler(JwtService ports.JWTservice) *AuthCheckHandler {
	return &AuthCheckHandler{JwtService}
}

// AuthCheck godoc
// @Summary      Проверка авторизации
// @Description  Приходит кука с токеном, высылается 200, если токен ещё валиден
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 	{object}  AuthCheckResponse
// @Failure      401    {object}  map[string]string
// @Router       /api/auth/check [get]
func (a *AuthCheckHandler) CheckAuth(c *gin.Context) {
	token, err := c.Cookie("jwt")
	if err != nil {
		authHeader := c.GetHeader("Authorisation")
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

	// c.Set("user", claims)
	delete(claims, "exp")
	c.JSON(http.StatusOK, AuthCheckResponse{
		IsAuthenticated: true,
		User:            claims,
	})
}

type AuthCheckResponse struct {
	IsAuthenticated bool                   `json:"isAuthenticated"`
	User            map[string]interface{} `json:"user"`
}
