package generic

import (
	"net/http"
	"strconv"

	genericPort "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GenericHandler[T, CreateDTO, ResponceDTO any] struct {
	repo genericPort.Repository[T]
}

func NewGenericHandler[T, C, R any](repo genericPort.Repository[T]) *GenericHandler[T, C, R] {
	return &GenericHandler[T, C, R]{repo: repo}
}

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

// ===========================================================UPDATE

func (h *GenericHandler[T, C, R]) Update(c *gin.Context) {
	id := c.Param("id")
	var fields map[string]interface{}
	if err := c.ShouldBindJSON(&fields); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(fields) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := h.repo.Update(id, fields); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
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

func (h *GenericHandler[T, C, R]) GetAllWhere(c *gin.Context) {
	var fieldsMap map[string]interface{}
	if err := c.ShouldBindJSON(&fieldsMap); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fields []string
	var values []interface{}

	for k, v := range fieldsMap {
		fields = append(fields, k)
		values = append(values, v)
	}

	var entities []T

	if err := h.repo.GetAllWhere(fields, values, &entities); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var resp []*R
	for _, e := range entities {
		resp = append(resp, mapper.MapDomainToDTO[T, R](&e))
	}

	c.JSON(http.StatusOK, resp)
}

// ===========================================================GetAll

func (h *GenericHandler[T, C, R]) GetAll(c *gin.Context) {

	var entities []T
	if err := h.repo.GetAll(&entities); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var resp []*R
	for _, e := range entities {
		resp = append(resp, mapper.MapDomainToDTO[T, R](&e))
	}
	c.JSON(http.StatusOK, resp)
}

// ===========================================================Delete

func (h *GenericHandler[T, C, R]) Delete(c *gin.Context) {
	idParam := c.Param("id")
	var id any

	// Попытка конвертации в int64
	if n, err := strconv.ParseInt(idParam, 10, 64); err == nil {
		id = n
	} else {
		// Если не число, оставляем как string (для UUID)
		id = idParam
	}

	var entity T
	if err := h.repo.Delete(id, &entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
