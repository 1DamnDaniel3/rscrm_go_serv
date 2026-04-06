package profilehandlers

import (
	"net/http"
	"strconv"

	profileucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/profile_ucs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetSelfProfileHandler struct {
	uc profileucs.IGetSelfProfileUC
}

func NewGetSelfProfileHandler(uc profileucs.IGetSelfProfileUC) *GetSelfProfileHandler {
	return &GetSelfProfileHandler{uc}
}

// Profile godoc
// @Summary      Логин
// @Description  Получить свой профиль пользователя по account_id
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        input  path     int64  true  "account_id по которому ищем профиль в пути маршрута"
// @Success      200	{object} dto.UserProfileResponseDTO
// @Header 		 200	{string} Set-Cookie "JWT-токен"
// @Failure      400    {object}  map[string]string
// @Router       /api/user_accounts/{id}/profile [get]
func (h *GetSelfProfileHandler) GetSelfProfile(c *gin.Context) {
	ctx := c.Request.Context()

	account_id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profiles := []entities.UserProfile{}

	if err := h.uc.Execute(ctx, account_id, &profiles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.UserProfile, dto.UserProfileResponseDTO](&profiles[0])

	c.JSON(http.StatusOK, resp)
}
