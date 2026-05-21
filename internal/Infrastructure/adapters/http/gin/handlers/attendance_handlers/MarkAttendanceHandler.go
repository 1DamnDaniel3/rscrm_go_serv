package attendancehandlers

import (
	"net/http"
	"strconv"

	attendanceucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/attendanceUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type MarkAttendanceHandler struct {
	uc attendanceucs.IMarkAttendanceUC
}

func NewMarkAttendanceHandler(
	uc attendanceucs.IMarkAttendanceUC,
) *MarkAttendanceHandler {
	return &MarkAttendanceHandler{uc}
}

// @Summary      MarkAttendances
// @Description  Отметка посещаемости absent <-> presence. Преподу доступны отметки только своих занятий.
// @Tags         Attendaces
// @Accept       json
// @Produce      json
// @Param        input  path     int64  true  "id посещения для отметки attendanceID"
// @Success 	 200  {object} dto.AttendanceResponseDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/attendances/mark/{attendanceID} [patch]
func (h *MarkAttendanceHandler) Mark(c *gin.Context) {
	ctx := c.Request.Context()

	attendanceID, err := strconv.ParseInt(c.Param("attendanceID"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.uc.Execute(ctx, attendanceID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Attendance, dto.AttendanceResponseDTO](updated)

	c.JSON(http.StatusOK, resp)
}
