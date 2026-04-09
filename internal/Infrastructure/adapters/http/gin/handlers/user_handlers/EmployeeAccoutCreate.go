package userhandlers

import (
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type CreateEmployeeAccountHandler struct {
	uc userUCs.ICreateEmployeeAccountUC
}

func NewCreateEmployeeAccountHandler(uc userUCs.ICreateEmployeeAccountUC) *CreateEmployeeAccountHandler {
	return &CreateEmployeeAccountHandler{uc}
}

// ================= DTO ===================

type CreateEmployeeAccountInputDTO struct {
	Account dto.UserAccountCreateDTO       `json:"account"`
	Profile dto.UserProfileCreateUpdateDTO `json:"profile"`
	Roles   []int64                        `json:"roles"`
}

type CreateEmployeeAccountResponse struct {
	Account dto.UserAccountResponseDTO `json:"account"`
	Profile dto.UserProfileResponseDTO `json:"profile"`
	Roles   []int64                    `json:"roles"`
}

// Employee godoc
// @Summary      Создать сотрудника
// @Description  Добавить сотрудника с ролями.
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param        input  body     CreateEmployeeAccountInputDTO  true  "Данные для регистрации сотрудника"
// @Success      201	{object} CreateEmployeeAccountResponse
// @Failure      400    {object}  map[string]string
// @Failure      500    {object}  map[string]string
// @Router       /api/user_accounts/createemployee [post]
func (h *CreateEmployeeAccountHandler) CreateEmployeeAccountHandler(c *gin.Context) {

	ctx := c.Request.Context()

	inputDto := &CreateEmployeeAccountInputDTO{}

	if err := c.ShouldBindJSON(inputDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	// === entities

	// account
	account := mapper.MapDTOToDomain[dto.UserAccountCreateDTO, entities.UserAccount](&inputDto.Account)

	// profile
	profile := mapper.MapDTOToDomain[dto.UserProfileCreateUpdateDTO, entities.UserProfile](&inputDto.Profile)

	// account_roles
	accountRoles := make([]entities.AccountRoles, len(inputDto.Roles))

	for i, v := range inputDto.Roles {
		accRole := entities.AccountRoles{
			Role_id: v,
		}
		accountRoles[i] = accRole
	}

	// EXECUTE UC

	if err := h.uc.Execute(ctx, account, profile, accountRoles); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	// response

	accountResp := mapper.MapDomainToDTO[entities.UserAccount, dto.UserAccountResponseDTO](account)
	profileResp := mapper.MapDomainToDTO[entities.UserProfile, dto.UserProfileResponseDTO](profile)

	resp := CreateEmployeeAccountResponse{
		Account: *accountResp,
		Profile: *profileResp,
		Roles:   inputDto.Roles,
	}

	c.JSON(http.StatusCreated, resp)
}
