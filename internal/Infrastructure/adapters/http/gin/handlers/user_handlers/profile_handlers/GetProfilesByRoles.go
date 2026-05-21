package profilehandlers

import (
	"net/http"

	profileucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/profile_ucs"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/mapper"
	"github.com/gin-gonic/gin"
)

type GetProfilesByRolesHandler struct {
	uc profileucs.IGetAllProfilesByRolesUC
}

func NewGetProfilesByRolesHandler(uc profileucs.IGetAllProfilesByRolesUC) *GetProfilesByRolesHandler {
	return &GetProfilesByRolesHandler{uc}
}

// ================================= DTO

type ProfilesByRolesResponseDTO struct {
	Data []dto.UserProfileResponseDTO `json:"data"`
}

// GetProfilesByRoles godoc
// @Summary      Профили сотрудников с их ролями
// @Description  Профили сотрудников школы для owner и абсолютно все для admin.
// @Description  Можно передать несколько параметров в path ?role=teacher?role=... получить все профили сотрудников с этими ролями.
// @Tags         Employee
// @Accept       json
// @Produce      json
// @Param role query string false "Роль для фильтра" Enums(admin, owner, manager, teacher, accountant, receptionist)
// @Success      200 	{object}  ProfilesByRolesResponseDTO
// @Failure      500    {object}  map[string]string
// @Router       /api/user_profile/profilesbyroles [get]
func (h *GetProfilesByRolesHandler) GetProfiles(c *gin.Context) {
	ctx := c.Request.Context()

	roles := c.QueryArray("role")

	profiles, err := h.uc.Execute(ctx, roles)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := &ProfilesByRolesResponseDTO{
		Data: make([]dto.UserProfileResponseDTO, len(profiles)),
	}

	for i := range profiles {
		profileDTO := mapper.MapDomainToDTO[entities.UserProfile, dto.UserProfileResponseDTO](&profiles[i])
		resp.Data[i] = *profileDTO
	}

	c.JSON(http.StatusOK, resp)
}
