package studenthandlers

import (
	"net/http"
	"strconv"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	bodtos "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/bo_dtos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type StudentClientsHandler struct {
	uc studentucs.IGetStudentClientsUC
}

func NewStudentClientsHandler(uc studentucs.IGetStudentClientsUC) *StudentClientsHandler {
	return &StudentClientsHandler{uc}
}

// @Summary      GetStudentClients
// @Description  Получить всех клиентов студента с метаданными (is_payer, relation)
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "ID студента"
// @Success      200  {object} GetStudentClientsResponseDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/students/{id}/clients [get]
func (h *StudentClientsHandler) GetStudentClients(c *gin.Context) {

	ctx := c.Request.Context()
	param_id := c.Param("id")

	student_id, err := strconv.ParseInt(param_id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studentClientsBO := []businessobjects.GetStudentClientsBO{}

	if err := h.uc.Execute(ctx, student_id, &studentClientsBO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := GetStudentClientsResponseDTO{
		Data: make([]bodtos.BoDTO_StudentClientsReponse, len(studentClientsBO)),
	}

	for i, bo := range studentClientsBO {
		clientR := mapper.MapDomainToDTO[entities.Client, dto.ClientResponseDTO](&bo.Client)
		clientDto := &bodtos.BoDTO_StudentClientsReponse{
			Relation_id:       bo.Relation_id,
			ClientResponseDTO: clientR,
			Is_payer:          bo.Is_payer,
			Relation:          bo.Relation,
		}
		resp.Data[i] = *clientDto
	}

	c.JSON(http.StatusOK, resp)
}

// =========== DTO ==========

type GetStudentClientsResponseDTO struct {
	Data []bodtos.BoDTO_StudentClientsReponse `json:"data"`
}
