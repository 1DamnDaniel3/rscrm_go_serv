package userprofilebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	profilehandlers "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/user_handlers/profile_handlers"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type UserProfileHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.UserProfile,
		dto.UserProfileCreateUpdateDTO,
		dto.UserProfileResponseDTO,
	]
	ProfilesByRolesHandler *profilehandlers.GetProfilesByRolesHandler
}

func NewUserProfileHandlerBuilder(
	uc *UserProfileUseCases,
) *UserProfileHandlerBuilder {

	return &UserProfileHandlerBuilder{
		CRUDHandler: generichandler.NewGenericHandler[
			entities.UserProfile,
			dto.UserProfileCreateUpdateDTO,
			dto.UserProfileResponseDTO,
		](uc.CRUD),

		// -=== ProfilesByRoles
		ProfilesByRolesHandler: profilehandlers.NewGetProfilesByRolesHandler(
			uc.GetProfilesByRoles,
		),
	}
}
