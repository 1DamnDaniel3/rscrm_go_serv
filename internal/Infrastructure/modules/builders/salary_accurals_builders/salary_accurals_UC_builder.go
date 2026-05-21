package salaryaccuralsbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type SalaryAccuralsUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.SalaryAccural]
}

func NewLessonSubscriptionsUseCases(
	tx services.Transaction,
	salaryAccuralsModule *modules.SalaryAccuralsModule,
) *SalaryAccuralsUseCases {

	return &SalaryAccuralsUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			salaryAccuralsModule.SalaryAccuralsRepo,
			salaryAccuralsModule.SalaryAccuralsPolicies.CRUD,
		),
	}
}
