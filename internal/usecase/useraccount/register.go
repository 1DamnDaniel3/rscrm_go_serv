package useraccount

import (
	"context"
	"time"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/domain/services"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/ports/generic"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/usecase/useraccount/dto"
)

type RegisterUseCase struct {
	tx             services.Transaction
	userRepo       generic.Repository[entities.UserAccount]
	profileRepo    generic.Repository[entities.UserProfile]
	schoolRepo     generic.Repository[entities.School]
	passwordHasher ports.PasswordHasher
}

func NewRegisterUseCase(
	tx services.Transaction,
	userRepo generic.Repository[entities.UserAccount],
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
			Name:      input.School.Name,
			City:      input.School.City,
			Phone:     input.School.Phone,
			Email:     input.School.Email,
			CreatedAt: now,
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
			FullName:  input.Profile.FullName,
			Phone:     input.Profile.Phone,
			BirthDate: input.Profile.Birthdate,
			AccountID: user.ID,
		}
		if err := uc.profileRepo.Create(profile); err != nil {
			return err
		}

		// Заполняем output
		output = dto.RegisterOutput{
			School: school.Name,
		}

		return nil
	})

	return output, err
}
