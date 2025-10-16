package user

import (
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type LoginUC struct {
	passwordHasher ports.PasswordHasher
	userRepo       entitiesrepos.UserAccountRepository
	JWTservice     ports.JWTservice
}

func NewLoginUseCase(
	hasher ports.PasswordHasher,
	repo entitiesrepos.UserAccountRepository,
	JWTservice ports.JWTservice) *LoginUC {
	return &LoginUC{hasher, repo, JWTservice}
}

func (uc *LoginUC) Execute(input *entities.UserAccount) (string, error) {
	account, err := uc.userRepo.GetByEmail(input.Email)
	if err != nil {
		return "", err
	}
	if !uc.passwordHasher.Compare(account.Password, input.Password) {
		return "", errors.New("wrong password")
	}
	// JWT SIGNING
	claims := map[string]interface{}{
		"user_id": account.ID,
		"email":   account.Email,
	}

	token, err := uc.JWTservice.Sign(claims)
	if err != nil {
		return "", err
	}
	return token, nil
}
