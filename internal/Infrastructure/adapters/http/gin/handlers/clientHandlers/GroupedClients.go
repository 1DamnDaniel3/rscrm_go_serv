package clienthandlers

import (
	"net/http"

	clientucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GroupedClientsHandler struct {
	uc clientucs.IGroupedClientsUC
}

func NewGroupedClientsHandler(uc clientucs.IGroupedClientsUC) *GroupedClientsHandler {
	return &GroupedClientsHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      GetGroupedClients
// @Description  Позволяет получить сгруппированных лидов школы
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        input  body     GetGroupedClientsInputDTO  true  "Фильтры: group_id и school_id(из токена авто)"
// @Success      200	{object} GetGroupedClientsOutputDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/leads/groupedclients [post]
func (h *GroupedClientsHandler) GetGroupedClients(c *gin.Context) {
	ctx := c.Request.Context()
	DTO := GetGroupedClientsInputDTO{}

	err := c.ShouldBindJSON(&DTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entitiesSlice := []entities.Client{}

	if err := h.uc.Execute(ctx, DTO.Group_id, &entitiesSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GetGroupedClientsOutputDTO{
		Data: make([]dto.ClientCreateUpdateDTO, 0, len(entitiesSlice)),
	}

	for _, client := range entitiesSlice {
		clientDto := mapper.MapDomainToDTO[entities.Client, dto.ClientCreateUpdateDTO](&client)
		output.Data = append(output.Data, *clientDto)
	}
	c.JSON(http.StatusOK, output)

}

// ============ DTO ============

type GetGroupedClientsInputDTO struct {
	Group_id int64 `json:"group_id"`
}

type GetGroupedClientsOutputDTO struct {
	Data []dto.ClientCreateUpdateDTO `json:"data"`
}
