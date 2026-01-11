package transactions

import (
	"net/http"
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/transactions/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterHandler struct {
	usecase userUCs.IRegisterUseCase
}

// Register godoc
// @Summary      Регистрация новой школы
// @Description  Создаёт школу, аккаунт владельца и профиль
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        input  body      dto.RegisterInput  true  "Данные для регистрации"
// @Success      200    {object}  dto.RegisterOutput
// @Failure      400    {object}  map[string]string
// @Router       /api/ownerschool/register [post]
func NewRegisterHandler(uc userUCs.IRegisterUseCase) *RegisterHandler {
	return &RegisterHandler{usecase: uc}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// маппинг dto в domain

	now := time.Now()

	// 1. Маппим школу
	school := &entities.School{
		Id:         uuid.New().String(),
		Name:       input.School.Name,
		City:       input.School.City,
		Phone:      input.School.Phone,
		Email:      input.School.Email,
		Created_at: now,
	}
	// 2. Маппим пользователя
	user := &entities.UserAccount{
		Email:      input.Account.Email,
		Password:   input.Account.Password,
		Created_at: now,
		School_id:  school.Id,
	}
	// 3. Создаём профиль
	profile := &entities.UserProfile{
		Full_name:  input.Profile.FullName,
		Phone:      input.Profile.Phone,
		Birthdate:  input.Profile.Birthdate,
		Account_id: user.Id,
	}
	accountRoles := &entities.AccountRoles{}

	outputEnity, err := h.usecase.Execute(c.Request.Context(), school, user, profile, accountRoles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// маппим ответ

	output := dto.RegisterOutput{
		School_id: outputEnity.Id,
	}

	c.JSON(http.StatusOK, output)
}
