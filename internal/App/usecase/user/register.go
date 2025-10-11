package user

import (
	"context"
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/user/dto"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
	"github.com/google/uuid"
)

type RegisterUseCase struct {
	tx             services.Transaction
	userRepo       ports.UserAccountRepository
	profileRepo    generic.Repository[entities.UserProfile]
	schoolRepo     generic.Repository[entities.School]
	passwordHasher ports.PasswordHasher
}

func NewRegisterUseCase(
	tx services.Transaction,
	userRepo ports.UserAccountRepository,
	profileRepo generic.Repository[entities.UserProfile],
	schoolRepo generic.Repository[entities.School],
	passwordHasher ports.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{tx, userRepo, profileRepo, schoolRepo, passwordHasher}
}

func (uc *RegisterUseCase) Execute(ctx context.Context, input dto.RegisterInput) (dto.RegisterOutput, error) {
	var output dto.RegisterOutput

	err := uc.tx.Do(ctx, func(txCtx context.Context) error {
		now := time.Now()

		// 1. Создаём школу
		school := &entities.School{
			ID:         uuid.New().String(),
			Name:       input.School.Name,
			City:       input.School.City,
			Phone:      input.School.Phone,
			Email:      input.School.Email,
			Created_at: now,
		}
		if err := uc.schoolRepo.Create(school); err != nil {
			return err
		}

		// 2. Создаём аккаунт
		hashedPassword, err := uc.passwordHasher.Hash(input.Account.Password)
		if err != nil {
			return err
		}

		user := &entities.UserAccount{
			Email:      input.Account.Email,
			Password:   hashedPassword,
			Role:       "owner",
			Created_at: now,
			School_id:  school.ID,
		}
		if err := uc.userRepo.Create(user); err != nil {
			return err
		}

		// 3. Создаём профиль
		profile := &entities.UserProfile{
			ID:         user.ID,
			Full_name:  input.Profile.FullName,
			Phone:      input.Profile.Phone,
			Birthdate:  input.Profile.Birthdate,
			Account_id: user.ID,
		}
		if err := uc.profileRepo.Create(profile); err != nil {
			return err
		}

		// Заполняем output
		output = dto.RegisterOutput{
			School_id: school.ID,
		}

		return nil
	})

	return output, err
}
