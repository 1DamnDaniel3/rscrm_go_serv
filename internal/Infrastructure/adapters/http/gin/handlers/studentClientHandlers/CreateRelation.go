package studentclienthandlers

import (
	"net/http"

	businessobjects "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/business_objects"
	studentclientsucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/studentClientsUCs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	bodtos "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/bo_dtos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type CreateStudentClientRelHandler struct {
	uc studentclientsucs.ICreateRelationUC
}

func NewCreateStudentClientRelHandler(uc studentclientsucs.ICreateRelationUC) *CreateStudentClientRelHandler {
	return &CreateStudentClientRelHandler{uc}
}

// GetGroupedLeads godoc
// @Summary      CreateRelation
// @Description  Нужен, чтобы создать связь, получив при этом в ответ StudentClients бизнес-объект StudentClients
// @Tags         StudentClient
// @Accept       json
// @Produce      json
// @Param        input  body     dto.StudentClientCreateUpdateDTO  true  "Обычный studentClient на вход"
// @Success      201	{object} bodtos.BoDTO_StudentClientsReponse
// @Failure      400    {object}  map[string]interface{}
// @Failure      500    {object}  map[string]interface{}
// @Router       /api/student-clients/createandgetBO [post]
func (h *CreateStudentClientRelHandler) CreateRel(c *gin.Context) {
	ctx := c.Request.Context()
	studCliDTO := dto.StudentClientCreateUpdateDTO{}

	if err := c.ShouldBindJSON(&studCliDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	studCli := mapper.MapDTOToDomain[dto.StudentClientCreateUpdateDTO, entities.StudentClient](&studCliDTO)

	bo := businessobjects.GetStudentClientsBO{}

	if err := h.uc.Execute(ctx, *studCli, &bo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := mapper.MapDomainToDTO[entities.Client, dto.ClientResponseDTO](&bo.Client)

	output := bodtos.BoDTO_StudentClientsReponse{
		Relation_id: bo.Relation_id,
		ClientAndGroups: bodtos.ClientAndGroups{
			ClientResponseDTO: client,
			Groups:            []dto.GroupResponseDTO{},
		},
		Is_payer: bo.Is_payer,
		Relation: bo.Relation,
	}

	c.JSON(http.StatusCreated, output)
}
