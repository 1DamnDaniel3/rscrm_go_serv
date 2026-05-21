package accountrolesucs

import (
	"context"

	accountrolespolicies "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/entities_policies/account_roles_policies"
	entitiesrepos "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/entities_repos"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/entities"
)

type AccountRolesCrudUC struct {
	UserAccountRepo entitiesrepos.UserAccountRepository
	// UserAccountPolicy useraccountpolicies.IUserAccountCrudPolicy

	AccountRolesRepo entitiesrepos.AccountRolesRepo
	CreatePolicy     accountrolespolicies.IAccountRolesCreatePolicy
	DeletePolicy     accountrolespolicies.IAccountRolesDeletePolicy
}

type IAccountRolesCrudUC interface {
	AssignRole(ctx context.Context, relation *entities.AccountRoles) error
	RemoveRole(ctx context.Context, relation *entities.AccountRoles) error
}

func NewAccountRolesCrudUC(
	UserAccountRepo entitiesrepos.UserAccountRepository,
	// UserAccountPolicy useraccountpolicies.IUserAccountCrudPolicy,

	AccountRolesRepo entitiesrepos.AccountRolesRepo,
	CreatePolicy accountrolespolicies.IAccountRolesCreatePolicy,
	DeletePolicy accountrolespolicies.IAccountRolesDeletePolicy,
) IAccountRolesCrudUC {
	return &AccountRolesCrudUC{
		UserAccountRepo: UserAccountRepo,
		// UserAccountPolicy: UserAccountPolicy,

		AccountRolesRepo: AccountRolesRepo,
		CreatePolicy:     CreatePolicy,
		DeletePolicy:     DeletePolicy,
	}
}

func (uc *AccountRolesCrudUC) AssignRole(ctx context.Context, relation *entities.AccountRoles) error {

	// acc_read_scope, err := uc.UserAccountPolicy.CanReadOne(ctx)
	// if err != nil {
	// 	return err
	// }

	// target user account
	emp_account := &entities.UserAccount{}
	if err := uc.UserAccountRepo.GetByID(ctx, relation.Account_id, emp_account, nil); err != nil {
		return err
	}

	scope, err := uc.CreatePolicy.CanAssignRole(ctx, relation.Role_id, relation.Account_id, emp_account.School_id)
	if err != nil {
		return err
	}

	if err := uc.AccountRolesRepo.Create(ctx, relation, scope); err != nil {
		return err
	}

	return nil
}

func (uc *AccountRolesCrudUC) RemoveRole(ctx context.Context, relation *entities.AccountRoles) error {

	relationMap := map[string]any{
		"account_id": relation.Account_id,
		"role_id":    relation.Role_id,
	}

	dbRelation, err := uc.AccountRolesRepo.FindRelation(ctx, relationMap, nil)
	if err != nil {
		return err
	}
	*relation = *dbRelation

	// target user account
	emp_account := &entities.UserAccount{}
	if err := uc.UserAccountRepo.GetByID(ctx, dbRelation.Account_id, emp_account, nil); err != nil {
		return err
	}

	del_scope, err := uc.DeletePolicy.CanRemoveRole(ctx, dbRelation.Role_id, dbRelation.Account_id, emp_account.School_id)
	if err != nil {
		return err
	}

	return uc.AccountRolesRepo.Delete(ctx, dbRelation.ID, dbRelation, del_scope)
}
