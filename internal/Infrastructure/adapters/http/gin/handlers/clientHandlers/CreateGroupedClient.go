package clienthandlers

import (
	"net/http"

	clientucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type CreateGroupedClientHandler struct {
	uc clientucs.ICreateGroupedClientUC
}

func NewCreateGroupedClientHandler(uc clientucs.ICreateGroupedClientUC) *CreateGroupedClientHandler {
	return &CreateGroupedClientHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      Createandgroup
// @Description  Позволяет получить сгруппированных лидов школы
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        input  body     СreateGroupedClientInputDto  true  "Фильтры: group_id и school_id(из токена авто)"
// @Success      200	{object} dto.ClientResponseDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/clients/createandgroup [post]
func (h *CreateGroupedClientHandler) CreateGroupedClient(c *gin.Context) {
	ctx := c.Request.Context()
	DTO := СreateGroupedClientInputDto{}

	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	client := mapper.MapDTOToDomain[dto.ClientCreateUpdateDTO, entities.Client](&DTO.Client)
	clientGroup := mapper.MapDTOToDomain[dto.ClientGroupCreateUpdateDTO, entities.ClientGroup](&DTO.ClientGroup)

	if err := h.uc.Execute(ctx, client, clientGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Client, dto.ClientResponseDTO](client)

	c.JSON(http.StatusOK, resp)
}

// ========== DTO ==========

type СreateGroupedClientInputDto struct {
	Client      dto.ClientCreateUpdateDTO      `json:"client"`
	ClientGroup dto.ClientGroupCreateUpdateDTO `json:"group"`
}
