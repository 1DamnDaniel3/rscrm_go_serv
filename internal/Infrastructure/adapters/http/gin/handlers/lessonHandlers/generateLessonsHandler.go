package lessonhandlers

import (
	"net/http"

	lessonshedulesucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/lessonsUCs/lessonShedulesUCs"
	"github.com/gin-gonic/gin"
)

type GenerateLessonsHandler struct {
	uc lessonshedulesucs.ICreateLessonsFromShceduleUC
}

func NewGenerateLessonsHandler(uc lessonshedulesucs.ICreateLessonsFromShceduleUC) *GenerateLessonsHandler {
	return &GenerateLessonsHandler{uc}
}

// @Summary      GenerateLessonsFromSchedule
// @Description  Сгенерировать занятия по шаблону расписания
// @Tags         Lessons
// @Accept       json
// @Produce      json
// @Success 	 200  {object} map[string]string
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/lessons/generate [get]
func (h *GenerateLessonsHandler) Generate(c *gin.Context) {
	ctx := c.Request.Context()

	if err := h.uc.Execute(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
