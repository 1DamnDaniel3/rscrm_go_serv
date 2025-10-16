package user

import (
	"net/http"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/gin-gonic/gin"
)

type CreateHandler struct {
	repo entitiesrepos.UserAccountRepository
}

func NewCreateHandler(repo entitiesrepos.UserAccountRepository) *CreateHandler {
	return &CreateHandler{repo: repo}
}

func (h *CreateHandler) CreateUserAccountHandler(c *gin.Context) {
	var DTO dto.UserAccountCreateDTO
	if err := c.ShouldBindJSON(&DTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// entity := mapper.MapDTOToDomain[dto.UserAccountCreateUpdateDTO, entities.UserAccount](&DTO)

}
