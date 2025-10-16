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
	uc *user.LoginUC
}

func NewLoginHandler(uc *user.LoginUC) *LoginHandler {
	return &LoginHandler{uc}
}

// Login godoc
// @Summary      Логин
// @Description  Вход стандарт email password, запись в httpOnly Cookies JWT
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input  body     dto.LoginDTO  true  "Данные для логина"
// @Success      200	{string} string "Ok. JWT установлен в httpOnly cookie"
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

	token, err := r.uc.Execute(entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	secure := os.Getenv("ENV") == "prod"

	c.SetCookie(
		"jwt",
		token,
		int(5*time.Hour),
		"/api",
		"localhost",
		secure,
		true,
	)

	c.JSON(http.StatusOK, gin.H{"message": "success"})

}
