package dancestylebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"

	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type DanceStyleUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.DanceStyle]
}

func NewDanceStyleUseCases(
	tx services.Transaction,
	danceStyleModule *modules.DanceStyleModule,

) *DanceStyleUseCases {

	return &DanceStyleUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			danceStyleModule.DanceStyleRepo,
			danceStyleModule.DanceStylePolicies.CRUD,
		),
	}
}
