package userUCs

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type RegisterUseCase struct {
	tx               services.Transaction
	userRepo         entitiesrepos.UserAccountRepository
	profileRepo      entitiesrepos.ProfileRepo
	schoolRepo       entitiesrepos.SchoolRepository
	accountRolesRepo entitiesrepos.AccountRolesRepo

	passwordHasher ports.PasswordHasher
}

type IRegisterUseCase interface {
	Execute(ctx context.Context,
		school *entities.School,
		userAccount *entities.UserAccount,
		userProfile *entities.UserProfile,
		accountRoles *entities.AccountRoles,
	) (*entities.School, error)
}

func NewRegisterUseCase(
	tx services.Transaction,
	userRepo entitiesrepos.UserAccountRepository,
	profileRepo entitiesrepos.ProfileRepo,
	schoolRepo entitiesrepos.SchoolRepository,
	accountRolesRepo entitiesrepos.AccountRolesRepo,

	passwordHasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{tx, userRepo, profileRepo, schoolRepo, accountRolesRepo, passwordHasher}
}

func (uc *RegisterUseCase) Execute(ctx context.Context,
	school *entities.School,
	userAccount *entities.UserAccount,
	userProfile *entities.UserProfile,
	accountRoles *entities.AccountRoles,
) (*entities.School, error) {

	err := uc.tx.Do(ctx, func(txCtx context.Context) error {
		// School Create
		if err := uc.schoolRepo.Register(txCtx, school); err != nil {
			return err
		}

		// UserAccount Create
		userAccount.School_id = school.Id
		if err := uc.userRepo.Register(txCtx, userAccount); err != nil {
			return err
		}

		userProfile.Id = userAccount.Id
		userProfile.Account_id = userAccount.Id
		userProfile.School_id = school.Id
		// UserProfile Create
		if err := uc.profileRepo.Register(txCtx, userProfile); err != nil {
			return err
		}

		// устанавливаем роль при регистрации - owner
		accountRoles.Role_id = 2
		accountRoles.Account_id = userAccount.Id
		accountRoles.School_id = school.Id
		if err := uc.accountRolesRepo.Register(txCtx, accountRoles); err != nil {
			return err
		}

		return nil
	})

	return school, err
}
