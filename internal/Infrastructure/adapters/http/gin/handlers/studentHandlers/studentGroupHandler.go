package studenthandlers

import (
	"net/http"
	"strconv"

	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
	studentucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type StudentGroupHandler struct {
	genericHandler *generichandler.GenericHandler[
		entities.StudentGroup,
		dto.StudentGroupCreateUpdateDTO,
		dto.StudentGroupResponseDTO,
	]
	createUC     studentucs.ICreateStudentUC
	getGroupedUC studentucs.IGroupedStudentsUC
}

func NewStudentGroupHandler(
	repo genericrepo.Repository[entities.StudentGroup],
	createUC studentucs.ICreateStudentUC,
	getGroupedUC studentucs.IGroupedStudentsUC,
) *StudentGroupHandler {
	return &StudentGroupHandler{
		generichandler.NewGenericHandler[
			entities.StudentGroup,
			dto.StudentGroupCreateUpdateDTO,
			dto.StudentGroupResponseDTO,
		](repo),
		createUC, getGroupedUC}
}

// ========= DTO =========

type СreateStudentInputDto struct {
	Student   dto.StudentCreateUpdateDTO      `json:"student"`
	StudGroup dto.StudentGroupCreateUpdateDTO `json:"group"`
}

type GroupedStudentsInputDTO struct {
	Group_id int64 `json:"group_id"`
}

type GroupedStudentsOutputDTO struct {
	Data []dto.StudentCreateUpdateDTO `json:"data"`
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
func (h *StudentGroupHandler) CreateStudent(c *gin.Context) {
	ctx := c.Request.Context()
	input := СreateStudentInputDto{}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := mapper.MapDTOToDomain[dto.StudentCreateUpdateDTO, entities.Student](&input.Student)
	studGroup := mapper.MapDTOToDomain[dto.StudentGroupCreateUpdateDTO, entities.StudentGroup](&input.StudGroup)

	if err := h.createUC.Execute(ctx, student, studGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.Student, dto.StudentResponseDTO](student)

	c.JSON(http.StatusOK, resp)
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
func (h *StudentGroupHandler) GetGroupedStudents(c *gin.Context) {
	inputDTO := GroupedStudentsInputDTO{}
	if err := c.ShouldBindJSON(&inputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()

	entitiesSlice := []entities.Student{}
	if err := h.getGroupedUC.Execute(ctx, inputDTO.Group_id, &entitiesSlice); err != nil {
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

// ================================================== CRUD ==========================================
// implement genericHandler

// CreateStudentGroup godoc
// @Summary Add student to group
// @Tags student-groups
// @Accept json
// @Produce json
// @Param body body dto.CreateStudentGroupDTO true "relation"
// @Success 201 {object} dto.StudentGroupResponseDTO
// @Router /student-groups [post]
func (h *StudentGroupHandler) CreateRelation(c *gin.Context) {

	studentID, _ := strconv.ParseInt(c.Param("studentId"), 10, 64)
	groupID, _ := strconv.ParseInt(c.Param("groupId"), 10, 64)

	dto := dto.StudentGroupCreateUpdateDTO{
		Student_id: studentID,
		Group_id:   groupID,
	}

	c.Set("dto", dto)

	h.genericHandler.Create(c)
}
