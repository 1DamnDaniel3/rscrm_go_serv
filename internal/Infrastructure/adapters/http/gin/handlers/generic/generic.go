package generic

import (
	"log"
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

type ResponceArrayDTO[T any] struct {
	Data *[]T `json:"data"`
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

	fields["id"] = id

	c.JSON(http.StatusOK, fields)
}

// ===========================================================GetByID
func (h *GenericHandler[T, C, R]) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	log.Println("PARAMETER ID IS ", idParam)
	var entity T

	if err := h.repo.GetByID(idParam, &entity); err != nil {
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

	var entities []T

	if err := h.repo.GetAllWhere(fieldsMap, &entities); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var data []*R
	for _, e := range entities {
		data = append(data, mapper.MapDomainToDTO[T, R](&e))
	}

	resp := &ResponceArrayDTO[*R]{
		Data: &data,
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

	var data []*R
	for _, e := range entities {
		data = append(data, mapper.MapDomainToDTO[T, R](&e))
	}

	resp := &ResponceArrayDTO[*R]{
		Data: &data,
	}

	c.JSON(http.StatusOK, resp)
}

// ===========================================================Delete

func (h *GenericHandler[T, C, R]) Delete(c *gin.Context) {
	idParam := c.Param("id")
	var entity T
	if err := h.repo.Delete(idParam, &entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"id": idParam})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})

}
