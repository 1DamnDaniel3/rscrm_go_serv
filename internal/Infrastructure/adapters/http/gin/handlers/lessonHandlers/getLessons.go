package lessonhandlers

import (
	"net/http"

	lessonsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/lessonsUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetLessonsHandler struct {
	uc lessonsucs.IGetLessonsUC
}

func NewGetLessonsHandler(uc lessonsucs.IGetLessonsUC) *GetLessonsHandler {
	return &GetLessonsHandler{uc}
}

// @Summary      FetchLessons
// @Description  Предоставляет актуальные Lesson, перед отдачей генерирует недостающие и удаляет старые
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Success 	 200  {object} GetLessonHandlerResDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/lessons/fetch [get]
func (h *GetLessonsHandler) GetLessons(c *gin.Context) {
	ctx := c.Request.Context()

	lessons, err := h.uc.Execute(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	resp := GetLessonHandlerResDTO{
		Data: make([]dto.LessonResponseDTO, len(lessons))}

	for i := 0; i < len(lessons); i++ {

		lesDTO := mapper.MapDomainToDTO[entities.Lesson, dto.LessonResponseDTO](&lessons[i])
		resp.Data[i] = *lesDTO

	}

	c.JSON(http.StatusOK, resp)

}

type GetLessonHandlerResDTO struct {
	Data []dto.LessonResponseDTO `json:"data"`
}
