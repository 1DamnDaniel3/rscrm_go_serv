package studentgroupHandlers

import (
	"net/http"
	"strconv"

	studentgroupUCs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentUCs/student_groupUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type StudentGroupHandler struct {
	createUC studentgroupUCs.ICreateStudentUC
	crudUCs  studentgroupUCs.IStudentGroupCRUDucs
}

func NewStudentGroupHandler(
	createUC studentgroupUCs.ICreateStudentUC,
	crudUCs studentgroupUCs.IStudentGroupCRUDucs,
) *StudentGroupHandler {
	return &StudentGroupHandler{

		createUC, crudUCs}
}

// ========= DTO =========

type СreateStudentInputDto struct {
	Student   dto.StudentCreateUpdateDTO      `json:"student"`
	StudGroup dto.StudentGroupCreateUpdateDTO `json:"group"`
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

// ================================================== CRUD ==========================================

// CreateStudentGroup godoc
// @Summary      Add student to group
// @Description  Adds a student to a group using their IDs from the path
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        studId   path      int  true  "ID of the student"
// @Param        groupId  path      int  true  "ID of the group"
// @Success      201 {object} dto.StudentGroupResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /students/{studId}/groups/{groupId} [post]
func (h *StudentGroupHandler) CreateRelation(c *gin.Context) {

	studentID, err := strconv.ParseInt(c.Param("studId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupID, err := strconv.ParseInt(c.Param("groupId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studGroup := entities.StudentGroup{
		Student_id: studentID,
		Group_id:   groupID,
	}

	if err := h.crudUCs.Create(c.Request.Context(), &studGroup); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[entities.StudentGroup, dto.StudentGroupResponseDTO](&studGroup)

	c.JSON(http.StatusOK, resp)
}

// CreateStudentGroup godoc
// @Summary      Add student to group
// @Description  Adds a student to a group using their IDs from the path
// @Tags         Students
// @Accept       json
// @Produce      json
// @Param        studId   path      int  true  "ID of the student"
// @Param        groupId  path      int  true  "ID of the group"
// @Success      201 {object} dto.StudentGroupResponseDTO
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /students/{studId}/groups/{groupId} [post]
func (h *StudentGroupHandler) DeleteRelation(c *gin.Context) {

	studentID, err := strconv.ParseInt(c.Param("studId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	groupID, err := strconv.ParseInt(c.Param("groupId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studGroup := entities.StudentGroup{
		Student_id: studentID,
		Group_id:   groupID,
	}

	// if err := h.crudUCs.Delete(c.Request.Context(), ); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	resp := mapper.MapDomainToDTO[entities.StudentGroup, dto.StudentGroupResponseDTO](&studGroup)

	c.JSON(http.StatusOK, resp)

}
