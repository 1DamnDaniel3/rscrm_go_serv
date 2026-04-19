package schoolpolicies

import (
	"context"
	"fmt"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	policyutils "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policy_utils"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/policytypes"
	"github.com/1DamnDaniel3/rscrm_go_serv/internal/Core/domain/valuetypes"
)

type SchoolCrudPolicy struct{}

type ISchoolCrudPolicy interface {
	crudpolicy.ICRUDPolicy
}

func NewSchoolCrudPolicy() ISchoolCrudPolicy {
	return &SchoolCrudPolicy{}
}

// ---====== methods ======---

func (p *SchoolCrudPolicy) CanCreate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Admin) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal: true,
	}, nil
}

// CanReadOne — проверяет доступ к одной школе
// Admin - любые школы
// Owner - только своя школа
func (p *SchoolCrudPolicy) CanReadOne(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	// админ может читать любые школы системы
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// владелец может читать только свою школу
	if policyutils.HasAnyRole(user, valuetypes.Owner) {
		return &policytypes.Scope{
			IsGlobal:  false,
			School_id: user.School_id,
		}, nil
	}

	// остальные роли не имеют доступа к школам
	return nil, fmt.Errorf("forbidden")
}

// CanReadAll — проверяет доступ к списку школ
// Использует ту же логику, что и CanReadOne
func (p *SchoolCrudPolicy) CanReadAll(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// CanReadAllWhere — проверяет доступ к выборке школ с фильтрами
// Логика идентична CanReadOne (ограничение по school_id или global)
func (p *SchoolCrudPolicy) CanReadAllWhere(ctx context.Context) (*policytypes.Scope, error) {
	return p.CanReadOne(ctx)
}

// CanUpdate — проверяет право обновлять школу
// Admin и Owner могут обновлять школу в рамках своего scope
func (p *SchoolCrudPolicy) CanUpdate(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Admin, valuetypes.Owner) {
		return nil, fmt.Errorf("forbidden")
	}

	// Admin → глобальный доступ
	if policyutils.HasAnyRole(user, valuetypes.Admin) {
		return &policytypes.Scope{
			IsGlobal: true,
		}, nil
	}

	// Owner → только своя школа
	return &policytypes.Scope{
		IsGlobal:  false,
		School_id: user.School_id,
	}, nil
}

// CanDelete — проверяет право удаления школы
// Только Admin может удалять школы
func (p *SchoolCrudPolicy) CanDelete(ctx context.Context) (*policytypes.Scope, error) {
	user, err := policyutils.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	if !policyutils.HasAnyRole(user, valuetypes.Admin) {
		return nil, fmt.Errorf("forbidden")
	}

	return &policytypes.Scope{
		IsGlobal: true,
	}, nil
}
