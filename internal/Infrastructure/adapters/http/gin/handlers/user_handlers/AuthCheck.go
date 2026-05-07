package userhandlers

import (
	"net/http"
	"strings"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/gin-gonic/gin"
)

type AuthCheckHandler struct {
	JwtService ports.JWTservice
	uc         genericcruduc.ICRUDUseCase[entities.UserAccount]
}

func NewAuthCheckHandler(
	JwtService ports.JWTservice,
	uc genericcruduc.ICRUDUseCase[entities.UserAccount],
) *AuthCheckHandler {
	return &AuthCheckHandler{JwtService, uc}
}

// AuthCheck godoc
// @Summary      Проверка авторизации
// @Description  Просто респект и уважуха если токен есть. Если нет то 401.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 	{object}  AuthCheckResponse
// @Failure      401    {object}  map[string]string
// @Router       /api/auth/check [get]
func (h *AuthCheckHandler) CheckAuth(c *gin.Context) {

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

	if _, err := h.JwtService.Verify(token); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	c.JSON(http.StatusOK, AuthCheckResponse{
		IsAuthenticated: true,
	})
}

type AuthCheckResponse struct {
	IsAuthenticated bool `json:"isAuthenticated"`
}
