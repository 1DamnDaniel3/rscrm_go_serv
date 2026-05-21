package accountroleshandlers

import (
	"net/http"
	"strconv"

	accountrolesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/account_roles_ucs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type AccountRolesCrudHandler struct {
	uc accountrolesucs.IAccountRolesCrudUC
}

func NewAccountRolesCrudHandler(
	uc accountrolesucs.IAccountRolesCrudUC,
) *AccountRolesCrudHandler {
	return &AccountRolesCrudHandler{uc}
}

// AssingRoles godoc
// @Summary      Добавить роль сотруднику
// @Description  Admin может раздавать роли всем, Owner только внутри школы. Себе нельзя выдать роль.
// @Tags         AccountRoles
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountRolesCreateUpdateDTO  true  "Данные для выдачи роли"
// @Success      201 	{object}  dto.AccountRolesResponseDTO
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/user_accounts/assignroles [post]
func (h *AccountRolesCrudHandler) AssingRoles(c *gin.Context) {

	ctx := c.Request.Context()
	accRolesDto := &dto.AccountRolesCreateUpdateDTO{}

	if err := c.ShouldBindJSON(accRolesDto); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	relation := mapper.MapDTOToDomain[dto.AccountRolesCreateUpdateDTO, entities.AccountRoles](accRolesDto)

	if err := h.uc.AssignRole(ctx, relation); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.AccountRoles, dto.AccountRolesResponseDTO](relation)

	c.JSON(http.StatusCreated, resp)
}

// RemoveRoles godoc
// @Summary      Убрать роль сотруднику
// @Description  Admin может убирать роли всем, Owner только внутри школы. Себе нельзя удалить роль. Роли admin и owner не стираются.
// @Tags         AccountRoles
// @Accept       json
// @Produce      json
// @Param        input  body      dto.AccountRolesCreateUpdateDTO  true  "Данные для выдачи роли"
// @Success      200 	{object}  dto.AccountRolesResponseDTO
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/user_account/{acc_id}/roles/{role_id} [delete]
func (h *AccountRolesCrudHandler) RemoveRoles(c *gin.Context) {

	ctx := c.Request.Context()

	acc_id, err := strconv.ParseInt(c.Param("acc_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	role_id, err := strconv.ParseInt(c.Param("role_id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	relation := &entities.AccountRoles{
		ID:         0,
		Account_id: acc_id,
		Role_id:    role_id,
		School_id:  "",
	}

	if err := h.uc.RemoveRole(ctx, relation); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.AccountRoles, dto.AccountRolesResponseDTO](relation)

	c.JSON(http.StatusOK, resp)
}
