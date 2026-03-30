package studenthandlers

import (
	"net/http"
	"strconv"

	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetStudentGroupsHandler struct {
	uc studentucs.IGetStudentGroupUC
}

func NewGetStudentGroupsHandler(uc studentucs.IGetStudentGroupUC) *GetStudentGroupsHandler {
	return &GetStudentGroupsHandler{uc}
}

// DTO
type GetClientGroupsResponceDTO struct {
	Data []dto.GroupResponseDTO `json:"data"`
}

// @Summary      GetGroupsByStudent
// @Description  Получить все группы студента
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "id - идентификатор ученика, чьи группы искать"
// @Success      200  {object} GetClientGroupsResponceDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/students/{id}/groups [get]
func (h *GetStudentGroupsHandler) GetClientGroups(c *gin.Context) {
	ctx := c.Request.Context()
	student_id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupSlice := []entities.Group{}
	if err := h.uc.Execute(ctx, student_id, &groupSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetClientGroupsResponceDTO{
		Data: make([]dto.GroupResponseDTO, len(groupSlice)),
	}

	for i, group := range groupSlice {
		groupDTO := mapper.MapDomainToDTO[entities.Group, dto.GroupResponseDTO](&group)
		resp.Data[i] = *groupDTO
	}

	c.JSON(http.StatusOK, resp)
}
