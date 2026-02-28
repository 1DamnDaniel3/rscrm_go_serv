package transactions

import (
	"net/http"
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
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
// @Param        input  body      RegisterInput  true  "Данные для регистрации"
// @Success      200    {object}  RegisterOutput
// @Failure      400    {object}  map[string]string
// @Router       /api/ownerschool/register [post]
func NewRegisterHandler(uc userUCs.IRegisterUseCase) *RegisterHandler {
	return &RegisterHandler{usecase: uc}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var input RegisterInput
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

	output := RegisterOutput{
		School_id: outputEnity.Id,
	}

	c.JSON(http.StatusOK, output)
}

// ============= DTO =============

// RegisterInput входные данные для регистрации
type RegisterInput struct {
	School  SchoolDTO  `json:"school"`
	Account AccountDTO `json:"account"`
	Profile ProfileDTO `json:"profile"`
}

// RegisterOutput данные, возвращаемые после успешной регистрации
type RegisterOutput struct {
	School_id string `json:"school_id" example:"bbeb26e7-7a3a-4bcf-8a70-338f362eabd1"`
}

// SchoolDTO данные о школе
type SchoolDTO struct {
	Name  string `json:"name" example:"Right Step"`
	City  string `json:"city" example:"Тимашевск"`
	Phone string `json:"phone" example:"+7-999-123-45-67"`
	Email string `json:"email" example:"popov@gmail.com"`
}

// AccountDTO данные для аккаунта
type AccountDTO struct {
	Email    string `json:"email" example:"popov@gmail.com"`
	Password string `json:"password" example:"secret"`
}

// ProfileDTO данные профиля пользователя
type ProfileDTO struct {
	FullName  string    `json:"full_name" example:"Артём Попов"`
	Phone     string    `json:"phone" example:"+7-999-123-45-67"`
	Birthdate time.Time `json:"birthdate" example:"1985-01-01T00:00:00Z"`
}
