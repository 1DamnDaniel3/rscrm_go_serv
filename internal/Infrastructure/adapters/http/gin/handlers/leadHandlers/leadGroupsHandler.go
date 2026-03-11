package leadhandlers

import (
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type LeadGroupsHandler struct {
	createUC     leadUCs.ICreateLeadUC
	getGroupedUC leadUCs.IGroupedLeadsUC
}

func NewLeadGroupsHandler(
	createUC leadUCs.ICreateLeadUC,
	getGroupedUC leadUCs.IGroupedLeadsUC) *LeadGroupsHandler {
	return &LeadGroupsHandler{createUC, getGroupedUC}
}

// ---===================== DTO ==========================---

type СreateLeadInputDto struct {
	Lead      dto.LeadCreateUpdateDTO      `json:"lead"`
	LeadGroup dto.LeadGroupCreateUpdateDTO `json:"group"`
}

// expecting from HTTP
type GroupedLeadInputDTO struct {
	Group_id int64 `json:"group_id"`
}

// answer
type GroupedLeadOutputDTO struct {
	Data []dto.LeadCreateUpdateDTO `json:"data"`
}

// LeadsAndGroups godoc
// @Summary      LeadsAndGroups
// @Description  Транзакция на создания лида с записью в lead_groups
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        input  body     СreateLeadInputDto  true  "Данные для создания лида и привязке к группе"
// @Success      200	{object} dto.LeadResponseDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/leads/createandgroup [post]
func (h *LeadGroupsHandler) CreateLead(c *gin.Context) {
	ctx := c.Request.Context()
	DTO := СreateLeadInputDto{}

	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lead := mapper.MapDTOToDomain[dto.LeadCreateUpdateDTO, entities.Lead](&DTO.Lead)
	leadGroup := mapper.MapDTOToDomain[dto.LeadGroupCreateUpdateDTO, entities.LeadGroup](&DTO.LeadGroup)

	if err := h.createUC.Execute(ctx, lead, leadGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Lead, dto.LeadResponseDTO](lead)

	c.JSON(http.StatusOK, resp)

}

// GetGroupedLeads godoc
// @Summary      GroupedLeads
// @Description  Позволяет получить сгруппированных лидов школы
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        input  body     GroupedLeadInputDTO  true  "Фильтры: group_id и school_id(из токена авто)"
// @Success      200	{object} GroupedLeadOutputDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/leads/groupedleads [post]
func (h *LeadGroupsHandler) GetGroupedLeads(c *gin.Context) {
	inputDTO := GroupedLeadInputDTO{}
	err := c.ShouldBindJSON(&inputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entitiesSlice := []entities.Lead{}

	ctx := c.Request.Context()

	if err := h.getGroupedUC.Execute(ctx, inputDTO.Group_id, &entitiesSlice); err != nil {
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
