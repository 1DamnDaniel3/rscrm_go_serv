package userprofilebuilders

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"

	generichandler "github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/adapters/http/gin/handlers/genericHandler"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Infrastructure/dto"
)

type UserProfileHandlerBuilder struct {
	CRUDHandler *generichandler.GenericHandler[
		entities.UserProfile,
		dto.UserProfileCreateUpdateDTO,
		dto.UserProfileResponseDTO,
	]
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
	}
}
