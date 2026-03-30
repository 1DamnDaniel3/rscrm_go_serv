package clienthandlers

import (
	"net/http"
	"strconv"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	clientstudentsUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/clientUCs/client_studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	bodtos "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/bo_dtos"
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

	studentSlice := []businessobjects.GetClientStudentsBO{}

	if err := h.uc.Execute(ctx, client_id, &studentSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	output := GetClientStudentsHandlerResponseDTO{
		Data: make([]bodtos.BoDTO_ClientStudentsReponse, len(studentSlice)),
	}

	for i, bo := range studentSlice {
		studR := mapper.MapDomainToDTO[entities.Student, dto.StudentResponseDTO](&bo.Student)
		studentDTO := bodtos.BoDTO_ClientStudentsReponse{
			Relation_id: bo.Relation_id,
			StudentsAndGroups: bodtos.StudentsAndGroups{
				StudentResponseDTO: studR,
				Groups:             []dto.GroupResponseDTO{},
			},
			Relation: bo.Relation,
		}
		output.Data[i] = studentDTO
	}

	c.JSON(http.StatusOK, output)
}

// ======================== DTO

type GetClientStudentsHandlerResponseDTO struct {
	Data []bodtos.BoDTO_ClientStudentsReponse `json:"data"`
}
