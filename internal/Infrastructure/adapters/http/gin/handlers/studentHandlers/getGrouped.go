package studenthandlers

import (
	"net/http"

	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetGroupedHandler struct {
	uc studentucs.IGroupedStudentsUC
}

func NewGetGroupedHandler(uc studentucs.IGroupedStudentsUC) *GetGroupedHandler {
	return &GetGroupedHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      GroupedStudents
// @Description  Позволяет получить сгруппированных учеников школы
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        input  body     GroupedStudentsInputDTO  true  "Фильтры: group_id и school_id(из токена авто)"
// @Success      200	{object} GroupedStudentsOutputDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/students/groupedstudents [post]
func (h *GetGroupedHandler) GetGroupedStudents(c *gin.Context) {
	inputDTO := GroupedStudentsInputDTO{}
	if err := c.ShouldBindJSON(&inputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	entitiesSlice := []entities.Student{}
	if err := h.uc.Execute(ctx, inputDTO.Group_id, &entitiesSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GroupedStudentsOutputDTO{
		Data: make([]dto.StudentCreateUpdateDTO, 0, len(entitiesSlice)),
	}

	for _, student := range entitiesSlice {
		StudDto := mapper.MapDomainToDTO[entities.Student, dto.StudentCreateUpdateDTO](&student)
		output.Data = append(output.Data, *StudDto)
	}

	c.JSON(http.StatusOK, output)

}

// ====== DTO

type GroupedStudentsInputDTO struct {
	Group_id int64 `json:"group_id"`
}

type GroupedStudentsOutputDTO struct {
	Data []dto.StudentCreateUpdateDTO `json:"data"`
}
