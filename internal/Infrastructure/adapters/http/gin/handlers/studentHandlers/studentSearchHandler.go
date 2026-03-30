package studenthandlers

import (
	"net/http"

	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type SearchStudentHandler struct {
	uc studentucs.ISearchStudentsUC
}

func NewSearchStudentHandler(uc studentucs.ISearchStudentsUC) *SearchStudentHandler {
	return &SearchStudentHandler{uc}
}

// @Summary      Search
// @Description  Через ?q= параметр получить students
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "q - QueryParameter"
// @Success      200  {object} StudentSearchHandlerDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/students/search [get]
func (h *SearchStudentHandler) Search(c *gin.Context) {

	ctx := c.Request.Context()
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	studentSlice := []entities.Student{}

	if err := h.uc.Execute(ctx, q, &studentSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	output := StudentSearchHandlerDTO{
		Data: make([]dto.StudentResponseDTO, len(studentSlice)),
	}

	for i, client := range studentSlice {
		studDto := mapper.MapDomainToDTO[entities.Student, dto.StudentResponseDTO](&client)
		output.Data[i] = *studDto
	}

	c.JSON(http.StatusOK, output)

}

// =============== DTO

type StudentSearchHandlerDTO struct {
	Data []dto.StudentResponseDTO `json:"data"`
}
