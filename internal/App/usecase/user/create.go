package user

import (
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
)

type CreateAccUC struct {
	passwordHasher ports.PasswordHasher
	userRepo       entitiesrepos.UserAccountRepository
}

func NewCreateAccountUseCase(
	repo entitiesrepos.UserAccountRepository,
	hasher ports.PasswordHasher) *CreateAccUC {
	return &CreateAccUC{userRepo: repo, passwordHasher: hasher}
}

// func (c *CreateAccUC) Execute (input )
