package generic

import (
	"gorm.io/gorm"
)

type GormRepository[T any] struct {
	db *gorm.DB
}

func (r *GormRepository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *GormRepository[T]) GetByID(id any, entity *T) error {
	return r.db.First(entity, "id=?", id).Error
}

func (r *GormRepository[T]) GetAllWhere(filters map[string]interface{}, entities *[]T) error {
	return r.db.Where(filters).Find(entities).Error
}

func (r *GormRepository[T]) GetAll(entities *[]T) error {
	return r.db.Find(entities).Error
}

func (r *GormRepository[T]) Update(id any, fields map[string]interface{}) error {
	if len(fields) == 0 {
		return nil
	}
	return r.db.Model(new(T)).Where("id = ?", id).Updates(fields).Error
}

func (r *GormRepository[T]) Delete(id any, entity *T) error {
	return r.db.Delete(entity, "id=?", id).Error
}

func NewGormRepository[T any](db *gorm.DB) *GormRepository[T] {
	return &GormRepository[T]{db: db}
}
