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

type GetAllAccountsWithRolesHandler struct {
	uc userUCs.IGetAllAccountsWithRolesUC
}

func NewGetAllAccountsWithRolesHandler(
	uc userUCs.IGetAllAccountsWithRolesUC,
) *GetAllAccountsWithRolesHandler {
	return &GetAllAccountsWithRolesHandler{uc}
}

// ========================== DTO =====

type AccsWithRolesResponseDTO struct {
	Data []bodtos.BoDTO_User `json:"data"`
}

// GetAccountsWithRoles godoc
// @Summary      Аккаунты сотрудников с их ролями
// @Description  Сотрудники школы с ролями для owner и абсолютно все для admin.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Success      200 	{object}  AccsWithRolesResponseDTO
// @Failure      500    {object}  map[string]string
// @Router       /api/user_accounts/allwithroles [get]
func (h *GetAllAccountsWithRolesHandler) GetAccountsWithRoles(c *gin.Context) {
	ctx := c.Request.Context()

	userBO, err := h.uc.Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := &AccsWithRolesResponseDTO{
		Data: make([]bodtos.BoDTO_User, len(userBO)),
	}

	for i := range userBO {
		userDTO := mapper.MapDomainToDTO[entities.UserAccount, dto.UserAccountResponseDTO](&userBO[i].UserAccount)
		userBODTO := &bodtos.BoDTO_User{
			User:  *userDTO,
			Roles: userBO[i].Roles,
		}

		resp.Data[i] = *userBODTO
	}

	c.JSON(http.StatusOK, resp)
}
