package clienthandlers

import (
	"net/http"

	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type ClientSearchHandler struct {
	repo entitiesrepos.ClientsQueryService
}

func NewClientSearchHandler(repo entitiesrepos.ClientsQueryService) *ClientSearchHandler {
	return &ClientSearchHandler{repo}
}

// @Summary      Search
// @Description  Через ?q= параметр получить clients
// @Tags         Clients
// @Accept       json
// @Produce      json
// @Param        id   path     int  true  "q - QueryParameter"
// @Success      200  {object} ClientSearchHandlerDTO
// @Failure      400  {object} map[string]interface{}
// @Failure      500  {object} map[string]interface{}
// @Router       /api/clients/search [get]
func (h *ClientSearchHandler) Search(c *gin.Context) {

	ctx := c.Request.Context()
	q := c.Query("q")
	if q == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query parameter 'q' is required"})
		return
	}

	clinetSlice := []entities.Client{}

	if err := h.repo.Search(ctx, q, &clinetSlice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	output := ClientSearchHandlerDTO{
		Data: make([]dto.ClientResponseDTO, len(clinetSlice)),
	}

	for i, client := range clinetSlice {
		clDto := mapper.MapDomainToDTO[entities.Client, dto.ClientResponseDTO](&client)
		output.Data[i] = *clDto
	}

	c.JSON(http.StatusOK, output)

}

// =============== DTO

type ClientSearchHandlerDTO struct {
	Data []dto.ClientResponseDTO `json:"data"`
}
