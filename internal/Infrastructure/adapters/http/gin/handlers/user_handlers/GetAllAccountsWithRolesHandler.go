package userhandlers

import (
	"net/http"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	bodtos "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/bo_dtos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetAllAccountsWithRolesHandler struct {
	uc         userUCs.IGetAllAccountsWithRolesUC
	ucFiltered userUCs.IGetAllAccountsWithRolesFilteredUC
}

func NewGetAllAccountsWithRolesHandler(
	uc userUCs.IGetAllAccountsWithRolesUC,
	ucFiltered userUCs.IGetAllAccountsWithRolesFilteredUC,
) *GetAllAccountsWithRolesHandler {
	return &GetAllAccountsWithRolesHandler{uc, ucFiltered}
}

// ========================== DTO =====

type AccsWithRolesResponseDTO struct {
	Data []bodtos.BoDTO_User `json:"data"`
}

// GetAccountsWithRoles godoc
// @Summary      Аккаунты сотрудников с их ролями
// @Description  Сотрудники школы с ролями для owner и абсолютно все для admin.
// @Description  Можно передать несколько параметров в path ?role=teacher&role=... получить всех сотрудников с этими ролями.
// @Description  Можно выполнить БЕЗ параметров (для owner и admin).
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param role query string false "Роль для фильтра" Enums(admin, owner, manager, teacher, accountant, receptionist)
// @Success      200 	{object}  AccsWithRolesResponseDTO
// @Failure      500    {object}  map[string]string
// @Router       /api/user_accounts/allwithroles [get]
func (h *GetAllAccountsWithRolesHandler) GetAccountsWithRoles(c *gin.Context) {
	ctx := c.Request.Context()

	roles := c.QueryArray("role")

	var userBO []*businessobjects.UserBO
	var err error

	if len(roles) > 0 {
		userBO, err = h.ucFiltered.Execute(ctx, roles)
	} else {
		userBO, err = h.uc.Execute(ctx)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
