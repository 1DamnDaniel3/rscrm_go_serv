package userhandlers

import (
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	bodtos "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/bo_dtos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetMeHandler struct {
	uc userUCs.IGetMeUC
}

func NewGetMeHandler(uc userUCs.IGetMeUC) *GetMeHandler {
	return &GetMeHandler{uc}
}

// ==== DTO

type GetMeResponse struct {
	User bodtos.BoDTO_User `json:"data"`
}

// GetMe godoc
// @Summary      Инфа аккаунта пользователя
// @Description  По токену ищется актуальная инфа пользователя из БД.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 	{object}  GetMeResponse
// @Failure      500    {object}  map[string]string
// @Router       /api/auth/check [get]
func (h *GetMeHandler) GetMe(c *gin.Context) {
	ctx := c.Request.Context()

	user, err := h.uc.Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	respDTO := bodtos.BoDTO_User{}

	userDTO := mapper.MapDTOToDomain[entities.UserAccount, dto.UserAccountResponseDTO](&user.UserAccount)

	respDTO.User = *userDTO
	respDTO.Roles = user.Roles

	resp := GetMeResponse{User: respDTO}

	c.JSON(http.StatusOK, resp)
}
