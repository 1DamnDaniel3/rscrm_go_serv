package salaryaccuralitemsbuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type SalaryAccuralItemsUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.SalaryAccrualItems]
}

func NewLessonSubscriptionsUseCases(
	tx services.Transaction,
	salaryAccuralItemsModule *modules.SalaryAccuralItemsModule,
) *SalaryAccuralItemsUseCases {

	return &SalaryAccuralItemsUseCases{
		// ================= CRUD =================
		CRUD: genericcruduc.NewCRUDUseCase(
			salaryAccuralItemsModule.SalaryAccuralItemsRepo,
			salaryAccuralItemsModule.SalaryAccuralItemsPolicies.CRUD,
		),
	}
}
