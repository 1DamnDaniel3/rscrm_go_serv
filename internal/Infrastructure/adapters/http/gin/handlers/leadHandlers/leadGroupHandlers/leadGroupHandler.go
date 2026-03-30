package leadgroupHandlers

import (
	"net/http"
	"strconv"

	leadgroupucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs/leadGroupUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type LeadGroupHandler struct {
	createUC leadgroupucs.ICreateLeadUC
	crudUCs  leadgroupucs.ILeadGroupCRUDucs
}

func NewLeadGroupHandler(
	createUC leadgroupucs.ICreateLeadUC,
	crudUCs leadgroupucs.ILeadGroupCRUDucs,
) *LeadGroupHandler {
	return &LeadGroupHandler{

		createUC, crudUCs}
}

// ========= DTO =========

type СreateLeadInputDto struct {
	Lead      dto.LeadCreateUpdateDTO      `json:"lead"`
	LeadGroup dto.LeadGroupCreateUpdateDTO `json:"group"`
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
func (h *LeadGroupHandler) CreateLead(c *gin.Context) {
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

// ================================================== CRUD ==========================================

// CreateLeadGroup godoc
// @Summary      Add lead to group
// @Description  Adds a lead to a group using their IDs from the path
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        leadId   path      int  true  "ID of the lead"
// @Param        groupId  path      int  true  "ID of the group"
// @Success      201 {object} dto.LeadGroupResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /students/{leadId}/groups/{groupId} [post]
func (h *LeadGroupHandler) CreateRelation(c *gin.Context) {

	leadID, err := strconv.ParseInt(c.Param("leadId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupID, err := strconv.ParseInt(c.Param("groupId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leadGroup := entities.LeadGroup{
		Lead_id:  leadID,
		Group_id: groupID,
	}

	if err := h.crudUCs.Create(c.Request.Context(), &leadGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.LeadGroup, dto.LeadGroupResponseDTO](&leadGroup)

	c.JSON(http.StatusOK, resp)
}

// DeleteLeadGroupRelation godoc
// @Summary      Delete relation lead_groups
// @Description  Удалить запись связи lead_groups
// @Tags         Leads
// @Accept       json
// @Produce      json
// @Param        leadId   path      int  true  "ID of the lead"
// @Param        groupId  path      int  true  "ID of the group"
// @Success      201 {object} dto.StudentGroupResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /leads/{leadId}/groups/{groupId} [delete]
func (h *LeadGroupHandler) DeleteRelation(c *gin.Context) {

	leadID, err := strconv.ParseInt(c.Param("leadId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupID, err := strconv.ParseInt(c.Param("groupId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leadGroup, err := h.crudUCs.Delete(c.Request.Context(), leadID, groupID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.LeadGroup, dto.LeadGroupResponseDTO](leadGroup)

	c.JSON(http.StatusOK, resp)
}
