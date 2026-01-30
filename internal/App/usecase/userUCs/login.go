package userUCs

import (
	"context"
	"errors"

	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type LoginUC struct {
	passwordHasher ports.PasswordHasher
	userRepo       entitiesrepos.UserAccountRepository
	accRolesRepo   entitiesrepos.AccountRolesRepo
	rolesRepo      entitiesrepos.RolesRepo
	JWTservice     ports.JWTservice
}

type ILoginUC interface {
	Execute(input *entities.UserAccount) (*entities.UserAccount, string, []string, error)
}

func NewLoginUseCase(
	hasher ports.PasswordHasher,
	repo entitiesrepos.UserAccountRepository,
	accRolesRepo entitiesrepos.AccountRolesRepo,
	rolesRepo entitiesrepos.RolesRepo,
	JWTservice ports.JWTservice) *LoginUC {
	return &LoginUC{hasher, repo, accRolesRepo, rolesRepo, JWTservice}
}

func (uc *LoginUC) Execute(input *entities.UserAccount) (*entities.UserAccount, string, []string, error) {
	account, err := uc.userRepo.GetByEmail(input.Email)
	if err != nil {
		return nil, "", make([]string, 0), err
	}

	if !uc.passwordHasher.Compare(account.Password, input.Password) {
		return nil, "", make([]string, 0), errors.New("wrong password")
	}

	accountRoleIds := []entities.AccountRoles{} // Получаем все роли аккаунта
	filters := make(map[string]interface{})
	filters["school_id"] = account.School_id
	filters["account_id"] = account.Id
	ctx := context.Background()
	if err := uc.accRolesRepo.GetAllWhere(ctx, filters, &accountRoleIds); err != nil {
		return nil, "", make([]string, 0), err
	}

	allRoles := &[]entities.Roles{} // Получаем вообще все роли
	if err := uc.rolesRepo.GetAll(ctx, allRoles); err != nil {
		return nil, "", make([]string, 0), err
	}
	allRolesMap := make(map[int64]string) //складываем в map
	for _, v := range *allRoles {
		allRolesMap[v.Id] = v.Role
	}

	accountRoleNames := make([]string, 0, 7) // Заполняем slice ролей
	for _, v := range accountRoleIds {
		accountRoleNames = append(accountRoleNames, allRolesMap[v.Role_id])
	}

	// JWT SIGNING
	claims := map[string]interface{}{
		"id":        account.Id,
		"roles":     accountRoleNames,
		"email":     account.Email,
		"school_id": account.School_id,
	}

	token, err := uc.JWTservice.Sign(claims)
	if err != nil {
		return nil, "", make([]string, 0), err
	}
	return account, token, accountRoleNames, nil
}
