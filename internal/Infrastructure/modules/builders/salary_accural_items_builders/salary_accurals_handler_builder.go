package salaryaccuralitemsbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type SalaryAccuralItemsHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.SalaryAccrualItems,
		dto.SalaryAccrualItemsCreateDTO,
		dto.SalaryAccrualItemsResponseDTO,
	]
}

func NewLessonSubscriptionsHandlerBuilder(
	uc *SalaryAccuralItemsUseCases,
) *SalaryAccuralItemsHandlerBuilder {

	return &SalaryAccuralItemsHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.SalaryAccrualItems,
			dto.SalaryAccrualItemsCreateDTO,
			dto.SalaryAccrualItemsResponseDTO,
		](uc.CRUD),
	}
}
