package leadhandlers

import (
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GroupedLeadsHandler struct {
	uc leadUCs.IGroupedLeadsUC
}

func NewGroupedLeadsHandler(uc leadUCs.IGroupedLeadsUC) *GroupedLeadsHandler {
	return &GroupedLeadsHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      GroupedLeads
// @Description  Позволяет получить сгруппированных лидов школы
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        input  body     GroupedLeadInputDTO  true  "Фильтры group_id и school_id"
// @Success      200	{object} GroupedLeadOutputDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/leads/groupedleads [post]
func (h *GroupedLeadsHandler) GetGroupedLeads(c *gin.Context) {
	inputDTO := GroupedLeadInputDTO{}
	err := c.ShouldBindJSON(&inputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entitiesSlice := []entities.Lead{}

	if err := h.uc.Execute(inputDTO.School_id, inputDTO.Group_id, &entitiesSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GroupedLeadOutputDTO{
		Data: make([]dto.LeadCreateUpdateDTO, 0, len(entitiesSlice)),
	}

	for _, lead := range entitiesSlice {
		dtolead := mapper.MapDomainToDTO[entities.Lead, dto.LeadCreateUpdateDTO](&lead)
		output.Data = append(output.Data, *dtolead)
	}

	c.JSON(http.StatusOK, output)

}

// expecting from HTTP
type GroupedLeadInputDTO struct {
	School_id string `json:"school_id"`
	Group_id  int64  `json:"group_id"`
}

// answer
type GroupedLeadOutputDTO struct {
	Data []dto.LeadCreateUpdateDTO `json:"data"`
}
