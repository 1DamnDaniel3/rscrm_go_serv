package school

import (
	"net/http"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/gin-gonic/gin"
)

type SchoolHandler struct {
	repo entitiesrepos.SchoolRepository
}

func NewSchoolHandler(repo entitiesrepos.SchoolRepository) *SchoolHandler {
	return &SchoolHandler{repo: repo}
}

func (h *SchoolHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	var entity entities.School
	h.repo.GetByID(id, &entity)

	c.JSON(http.StatusOK, entity)
}
