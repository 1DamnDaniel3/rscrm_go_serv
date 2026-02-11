package studenthandlers

import (
	"fmt"
	"net/http"

	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type CreateStudentHandler struct {
	uc studentucs.ICreateStudentUC
}

func NewCreateStudentHandler(uc studentucs.ICreateStudentUC) *CreateStudentHandler {
	return &CreateStudentHandler{uc}
}

// LeadsAndGroups godoc
// @Summary      StudentsAndGroups
// @Description  Транзакция на создания ученика с записью в student_groups
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        input  body     СreateStudentInputDto  true  "Данные для создания студента и привязке к группе"
// @Success      200	{object} dto.StudentResponseDTO
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/students/createandgroup [post]
func (h *CreateStudentHandler) CreateStudent(c *gin.Context) {
	ctx := c.Request.Context()
	input := СreateStudentInputDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("==========INPUT: ", input)

	student := mapper.MapDTOToDomain[dto.StudentCreateUpdateDTO, entities.Student](&input.Student)
	studGroup := mapper.MapDTOToDomain[dto.StudentGroupCreateUpdateDTO, entities.StudentGroup](&input.StudGroup)

	fmt.Println("==========STUDENT: ", input)
	fmt.Println("==========STUD_GROUP: ", input)

	if err := h.uc.Execute(ctx, student, studGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Student, dto.StudentResponseDTO](student)

	c.JSON(http.StatusOK, resp)
}

// ========= DTO =========

type СreateStudentInputDto struct {
	Student   dto.StudentCreateUpdateDTO      `json:"student"`
	StudGroup dto.StudentGroupCreateUpdateDTO `json:"group"`
}
