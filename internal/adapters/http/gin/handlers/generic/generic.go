package generic

import (
	"net/http"
	"strconv"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/mapper"
	genericPort "github.com/1DamnDaniel3/rscrm_go_serv/internal/ports/generic"
	"github.com/gin-gonic/gin"
)

type GenericHandler[T, CreateDTO, ResponceDTO any] struct {
	repo genericPort.Repository[T]
}

func NewGenericHandler[T, C, R any](repo genericPort.Repository[T]) *GenericHandler[T, C, R] {
	return &GenericHandler[T, C, R]{repo: repo}
}

// @Summary Register new user
// @Description Create a new user account
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserAccountCreateUpdateDTO true "User Data"
// @Success 201 {object} dto.UserAccountResponseDTO
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user_accounts/register [post]
// ===========================================================CREATE
func (h *GenericHandler[T, C, R]) Create(c *gin.Context) {
	var dto C
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := mapper.MapDTOToDomain[C, T](&dto)

	if err := h.repo.Create(entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := mapper.MapDomainToDTO[T, R](entity)
	c.JSON(http.StatusCreated, resp)
}

// ===========================================================GetByID
func (h *GenericHandler[T, C, R]) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var entity T

	if err := h.repo.GetByID(id, &entity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.MapDomainToDTO[T, R](&entity)
	c.JSON(http.StatusOK, resp)
}

// ===========================================================GetAllWhere
