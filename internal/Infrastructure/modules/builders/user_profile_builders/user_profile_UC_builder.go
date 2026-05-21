package userprofilebuilders

import (
	profileucs "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/entitiesUCs/userUCs/profile_ucs"
	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/modules"
)

type UserProfileUseCases struct {
	CRUD               genericcruduc.ICRUDUseCase[entities.UserProfile]
	GetProfilesByRoles profileucs.IGetAllProfilesByRolesUC
}

func NewUserProfileUseCasesBuilder(
	profileModule *modules.ProfileModule,
	accountModule *modules.AccountModule,
) *UserProfileUseCases {

	return &UserProfileUseCases{
		CRUD: genericcruduc.NewCRUDUseCase(
			profileModule.ProfileRepo,
			profileModule.ProfilePolicies.CRUD,
		),

		GetProfilesByRoles: profileucs.NewGetAllProfilesByRolesUC(
			profileModule.ProfileRepo,
			accountModule.AccountPolicies.ReadPolicy,
		),
	}
}
