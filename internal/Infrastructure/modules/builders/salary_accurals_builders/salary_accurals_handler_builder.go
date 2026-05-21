package salaryaccuralsbuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type SalaryAccuralsHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.SalaryAccural,
		dto.SalaryAccuracyCreateDTO,
		dto.SalaryAccuracyResponseDTO,
	]
}

func NewLessonSubscriptionsHandlerBuilder(
	uc *SalaryAccuralsUseCases,
) *SalaryAccuralsHandlerBuilder {

	return &SalaryAccuralsHandlerBuilder{
		// ================= CRUD =================
		CRUDHandler: generichandler.NewGenericHandler[
			entities.SalaryAccural,
			dto.SalaryAccuracyCreateDTO,
			dto.SalaryAccuracyResponseDTO,
		](uc.CRUD),
	}
}
