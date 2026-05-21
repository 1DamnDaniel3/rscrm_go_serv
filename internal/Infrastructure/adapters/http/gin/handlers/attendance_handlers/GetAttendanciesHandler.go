package attendancehandlers

import (
	"net/http"

	attendanceucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/attendanceUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GenerateAttendanciesHandler struct {
	uc attendanceucs.IGetAttendanciesUC
}

func NewGenerateAttendanciesHandler(
	uc attendanceucs.IGenerateAttendanceUC,
) *GenerateAttendanciesHandler {
	return &GenerateAttendanciesHandler{uc}
}

// ============================ DTO ====================

type GenerateAttendanceInputDTO struct {
	Group_id  int64 `json:"group_id" binding:"required"`
	Lesson_id int64 `json:"lesson_id" binding:"required"`
}

type GenerateAttendanceOutputDTO struct {
	Data []dto.AttendanceResponseDTO `json:"data"`
}

// @Summary      GenerateAttendancies
// @Description  Посещаемость по group_id и lesson_id. Генерирует посещаемость в момент вызова.
// @Tags         Attendaces
// @Accept       json
// @Produce      json
// @Param        input  body     GenerateAttendanceInputDTO  true  "Данные для генерации"
// @Success 	 200  {object} GenerateAttendanceOutputDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/attendances/generate [post]
func (h *GenerateAttendanciesHandler) Generate(c *gin.Context) {

	ctx := c.Request.Context()

	input := &GenerateAttendanceInputDTO{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attendancies, err := h.uc.Execute(ctx, input.Group_id, input.Lesson_id)
	if err != nil {
		if err.Error() == "empty group" {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": "Cannot generate attendances. Students group is empty"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := &GenerateAttendanceOutputDTO{
		Data: make([]dto.AttendanceResponseDTO, len(attendancies)),
	}

	for i := range attendancies {
		attDTO := mapper.MapDomainToDTO[entities.Attendance, dto.AttendanceResponseDTO](&attendancies[i])
		resp.Data[i] = *attDTO
	}

	c.JSON(http.StatusOK, resp)

}
