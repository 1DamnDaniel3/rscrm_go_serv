package userUCs

import (
	"context"
	"fmt"

	genericcruduc "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/usecase/generic_crud_uc"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/services"
)

type CreateEmployeeAccountUC struct {
	tx                 services.Transaction
	accountCrudUC      genericcruduc.ICRUDUseCase[entities.UserAccount]
	profileCrudUC      genericcruduc.ICRUDUseCase[entities.UserProfile]
	accountRolesCrudUC genericcruduc.ICRUDUseCase[entities.AccountRoles]
}

type ICreateEmployeeAccountUC interface {
	Execute(ctx context.Context,
		account *entities.UserAccount,
		profile *entities.UserProfile,
		roles []entities.AccountRoles) error
}

func NewCreateEmployeeAccountUC(
	tx services.Transaction,
	accountCrudUC genericcruduc.ICRUDUseCase[entities.UserAccount],
	profileCrudUC genericcruduc.ICRUDUseCase[entities.UserProfile],
	accountRolesCrudUC genericcruduc.ICRUDUseCase[entities.AccountRoles]) ICreateEmployeeAccountUC {
	return &CreateEmployeeAccountUC{
		tx:                 tx,
		accountCrudUC:      accountCrudUC,
		profileCrudUC:      profileCrudUC,
		accountRolesCrudUC: accountRolesCrudUC,
	}
}

func (uc *CreateEmployeeAccountUC) Execute(
	ctx context.Context,
	account *entities.UserAccount,
	profile *entities.UserProfile,
	roles []entities.AccountRoles,
) error {

	if err := uc.tx.Do(ctx, func(txCtx context.Context) error {

		if len(roles) == 0 {
			return fmt.Errorf("Missing roles")
		}
		for _, role := range roles {
			if role.Role_id == 1 || role.Role_id == 2 {
				return fmt.Errorf("You cannot create new admin or owner")
			}
		}

		// Account
		if err := uc.accountCrudUC.Create(txCtx, account); err != nil {
			return err
		}

		// Profile
		profile.Id = account.Id
		profile.Account_id = account.Id
		if err := uc.profileCrudUC.Create(txCtx, profile); err != nil {
			return err
		}

		for i := range roles {
			roles[i].Account_id = account.Id
		}

		// Roles
		if err := uc.accountRolesCrudUC.CreateMany(txCtx, &roles); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
