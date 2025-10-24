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

func (uc *LoginUC) Execute(input *entities.UserAccount) (*entities.UserAccount, string, error) {
	account, err := uc.userRepo.GetByEmail(input.Email)
	if err != nil {
		return nil, "", err
	}

	if !uc.passwordHasher.Compare(account.Password, input.Password) {
		return nil, "", errors.New("wrong password")
	}
	// JWT SIGNING
	claims := map[string]interface{}{
		"id":        account.ID,
		"role":      account.Role,
		"email":     account.Email,
		"school_id": account.School_id,
	}

	token, err := uc.JWTservice.Sign(claims)
	if err != nil {
		return nil, "", err
	}
	return account, token, nil
}
