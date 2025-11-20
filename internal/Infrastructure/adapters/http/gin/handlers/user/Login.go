package user

import (
	"net/http"
	"os"
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/user"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	uc user.ILoginUC
}

func NewLoginHandler(uc user.ILoginUC) *LoginHandler {
	return &LoginHandler{uc}
}

// Login godoc
// @Summary      Логин
// @Description  Вход стандарт email password, запись в httpOnly Cookies JWT
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input  body     dto.LoginDTO  true  "Данные для логина"
// @Success      200	{object} LoginResponse
// @Header 		 200	{string} Set-Cookie "JWT-токен"
// @Failure      400    {object}  map[string]string
// @Router       /api/useraccounts/login [post]
func (r *LoginHandler) Login(c *gin.Context) {
	var DTO dto.LoginDTO
	err := c.ShouldBindJSON(&DTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	entity := mapper.MapDTOToDomain[dto.LoginDTO, entities.UserAccount](&DTO)

	account, token, err := r.uc.Execute(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	secure := os.Getenv("ENV") == "prod"
	var host string
	if os.Getenv("ENV") == "prod" {
		host = os.Getenv("HOST")
	} else {
		host = "localhost"
	}

	c.SetCookie(
		"jwt",
		token,
		int(5*time.Hour),
		"/api",
		host,
		secure,
		true,
	)

	response := mapper.MapDomainToDTO[entities.UserAccount, dto.UserAccountResponseDTO](account)

	c.JSON(http.StatusOK, LoginResponse{
		Message: "success",
		User:    response,
	})

}

// === DTO ===

type LoginResponse struct {
	Message string                      `json:"message" example:"success"`
	User    *dto.UserAccountResponseDTO `json:"user"`
}
