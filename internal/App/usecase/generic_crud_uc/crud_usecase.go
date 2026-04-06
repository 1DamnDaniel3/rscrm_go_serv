package genericcruduc

import (
	"context"

	crudpolicy "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/policies/crud_policy"
	genericrepo "github.com/1DamnDaniel3/rscrm_go_serv/internal/App/ports/genericRepo"
)

type CRUDUseCase[T any] struct {
	repo   genericrepo.Repository[T]
	policy crudpolicy.ICRUDPolicy
}

type ICRUDUseCase[T any] interface {
	Create(ctx context.Context, entity *T) error
	CreateMany(ctx context.Context, entities *[]T) error
	GetByID(ctx context.Context, id any, entity *T) error
	GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T) error
	GetAll(ctx context.Context, entities *[]T) error
	Update(ctx context.Context, id any, fields map[string]interface{}) error
	Delete(ctx context.Context, id any, entity *T) error
	FindRelation(ctx context.Context, relationMap map[string]any) (*T, error)
}

func NewCRUDUseCase[T any](
	repo genericrepo.Repository[T],
	policy crudpolicy.ICRUDPolicy) ICRUDUseCase[T] {
	return &CRUDUseCase[T]{repo: repo, policy: policy}
}

// ====================================== METHODS =================================

// ====================================== CREATE
func (uc *CRUDUseCase[T]) Create(ctx context.Context, entity *T) error {
	if err := uc.policy.CanCreate(ctx); err != nil {
		return err
	}
	return uc.repo.Create(ctx, entity)
}

// ====================================== CREATE MANY
func (uc *CRUDUseCase[T]) CreateMany(ctx context.Context, entities *[]T) error {
	if err := uc.policy.CanCreate(ctx); err != nil {
		return err
	}
	return uc.repo.CreateMany(ctx, entities)
}

// ====================================== GET BY ID
func (uc *CRUDUseCase[T]) GetByID(ctx context.Context, id any, entity *T) error {
	if err := uc.policy.CanRead(ctx); err != nil {
		return err
	}
	return uc.repo.GetByID(ctx, id, entity)
}

// ====================================== GET ALL WHERE
func (uc *CRUDUseCase[T]) GetAllWhere(ctx context.Context, filters map[string]interface{}, entities *[]T) error {
	if err := uc.policy.CanRead(ctx); err != nil {
		return err
	}
	return uc.repo.GetAllWhere(ctx, filters, entities)
}

// ====================================== GET ALL
func (uc *CRUDUseCase[T]) GetAll(ctx context.Context, entities *[]T) error {
	if err := uc.policy.CanRead(ctx); err != nil {
		return err
	}
	return uc.repo.GetAll(ctx, entities)
}

// ====================================== UPDATE
func (uc *CRUDUseCase[T]) Update(ctx context.Context, id any, fields map[string]interface{}) error {
	if err := uc.policy.CanUpdate(ctx); err != nil {
		return err
	}
	return uc.repo.Update(ctx, id, fields)
}

// ====================================== DELETE
func (uc *CRUDUseCase[T]) Delete(ctx context.Context, id any, entity *T) error {
	if err := uc.policy.CanDelete(ctx); err != nil {
		return err
	}
	return uc.repo.Delete(ctx, id, entity)
}

// ====================================== FIND RELATION
func (uc *CRUDUseCase[T]) FindRelation(ctx context.Context, relationMap map[string]any) (*T, error) {
	if err := uc.policy.CanRead(ctx); err != nil {
		return nil, err
	}
	return uc.repo.FindRelation(ctx, relationMap)
}
