package clienthandlers

import (
	"net/http"
	"strconv"

	clientstudentsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs/client_studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetClientStudentsHandler struct {
	uc clientstudentsUCs.IGetClientStudentsUC
}

func NewGetClientStudentsHandler(uc clientstudentsUCs.IGetClientStudentsUC) *GetClientStudentsHandler {
	return &GetClientStudentsHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      GetClientStudents
// @Description  Позволяет получить всех студентов клиента
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        input  path     int64  true  "client_id"
// @Success      200	{object} GetClientStudentsHandlerResponseDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/clients/{id}/students [get]
func (h *GetClientStudentsHandler) GetClientStudents(c *gin.Context) {
	ctx := c.Request.Context()
	param_id := c.Param("id")
	client_id, err := strconv.ParseInt(param_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentSlice := []entities.Student{}

	if err := h.uc.Execute(ctx, client_id, &studentSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GetClientStudentsHandlerResponseDTO{
		Data: make([]dto.StudentResponseDTO, len(studentSlice)),
	}

	for i, student := range studentSlice {
		studDTO := mapper.MapDomainToDTO[entities.Student, dto.StudentResponseDTO](&student)
		output.Data[i] = *studDTO
	}

	c.JSON(http.StatusOK, output)
}

// ======================== DTO

type GetClientStudentsHandlerResponseDTO struct {
	Data []dto.StudentResponseDTO `json:"data"`
}
