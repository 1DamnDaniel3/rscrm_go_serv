package user

import (
	"context"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type RegisterUseCase struct {
	tx             services.Transaction
	userRepo       entitiesrepos.UserAccountRepository
	profileRepo    entitiesrepos.ProfileRepo
	schoolRepo     entitiesrepos.SchoolRepository
	passwordHasher ports.PasswordHasher
}

func NewRegisterUseCase(
	tx services.Transaction,
	userRepo entitiesrepos.UserAccountRepository,
	profileRepo entitiesrepos.ProfileRepo,
	schoolRepo entitiesrepos.SchoolRepository,
	passwordHasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{tx, userRepo, profileRepo, schoolRepo, passwordHasher}
}

func (uc *RegisterUseCase) Execute(ctx context.Context,
	school *entities.School,
	userAccount *entities.UserAccount,
	userProfile *entities.UserProfile) (*entities.School, error) {

	err := uc.tx.Do(ctx, func(txCtx context.Context) error {
		// School Create
		if err := uc.schoolRepo.Register(txCtx, school); err != nil {
			return err
		}

		// UserAccount Create
		if err := uc.userRepo.Register(txCtx, userAccount); err != nil {
			return err
		}

		userProfile.ID = userAccount.ID
		userProfile.Account_id = userAccount.ID
		// UserProfile Create
		if err := uc.profileRepo.Register(txCtx, userProfile); err != nil {
			return err
		}

		return nil
	})

	return school, err
}
