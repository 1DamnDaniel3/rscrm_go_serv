package userprofilebuilders

import (
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type UserProfileUseCases struct {
	CRUD genericcruduc.ICRUDUseCase[entities.UserProfile]
}

func NewUserProfileUseCasesBuilder(
	profileModule *modules.ProfileModule,
) *UserProfileUseCases {

	return &UserProfileUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			profileModule.ProfileRepo,
			profileModule.ProfilePolicies.CRUD,
		),
	}
}
