package transactions

import (
	"net/http"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/usecase/user"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/usecase/user/dto"
	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	usecase *user.RegisterUseCase
}

func NewRegisterHandler(uc *user.RegisterUseCase) *RegisterHandler {
	return &RegisterHandler{usecase: uc}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var input dto.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.usecase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}
