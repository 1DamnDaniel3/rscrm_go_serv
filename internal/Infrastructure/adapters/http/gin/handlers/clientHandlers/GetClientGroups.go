package clienthandlers

import (
	"net/http"
	"strconv"

	clientgroupsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs/client_groupUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetClientGroupsHandler struct {
	uc clientgroupsUCs.IGetClientGroupUC
}

func NewGetClientGroupsHandler(uc clientgroupsUCs.IGetClientGroupUC) *GetClientGroupsHandler {
	return &GetClientGroupsHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      GetClientGroups
// @Description  Позволяет получить все группы, в которых состоит клиент
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        input  path     int64  true  "client_id"
// @Success      200	{object} GetGroupsRequestDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/clients/{id}/groups [get]
func (h *GetClientGroupsHandler) GetGroups(c *gin.Context) {
	ctx := c.Request.Context()
	param_id := c.Param("id")
	client_id, err := strconv.ParseInt(param_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	groupSlice := []entities.Group{}

	if err := h.uc.Execute(ctx, client_id, &groupSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GetGroupsRequestDTO{
		Data: make([]dto.GroupResponseDTO, len(groupSlice)),
	}

	for i, group := range groupSlice {
		groupDto := mapper.MapDomainToDTO[entities.Group, dto.GroupResponseDTO](&group)
		output.Data[i] = *groupDto
	}

	c.JSON(http.StatusOK, output)

}

// ====================== DTO

type GetGroupsRequestDTO struct {
	Data []dto.GroupResponseDTO `json:"data"`
}
