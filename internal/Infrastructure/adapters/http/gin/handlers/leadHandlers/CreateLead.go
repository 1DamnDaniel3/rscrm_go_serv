package leadhandlers

import (
	"fmt"
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/leadUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type CreateLeadHandler struct {
	uc leadUCs.ICreateLeadUC
}

func NewCreateLeadHandler(uc leadUCs.ICreateLeadUC) *CreateLeadHandler {
	return &CreateLeadHandler{uc: uc}
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
func (h *CreateLeadHandler) CreateLead(c *gin.Context) {
	ctx := c.Request.Context()
	DTO := СreateLeadInputDto{}

	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(DTO)
	lead := mapper.MapDTOToDomain[dto.LeadCreateUpdateDTO, entities.Lead](&DTO.Lead)
	leadGroup := mapper.MapDTOToDomain[dto.LeadGroupCreateUpdateDTO, entities.LeadGroup](&DTO.LeadGroup)

	if err := h.uc.Execute(ctx, lead, leadGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Lead, dto.LeadResponseDTO](lead)

	c.JSON(http.StatusOK, resp)

}

// ---===================== DTO ==========================---

type СreateLeadInputDto struct {
	Lead      dto.LeadCreateUpdateDTO      `json:"lead"`
	LeadGroup dto.LeadGroupCreateUpdateDTO `json:"group"`
}
