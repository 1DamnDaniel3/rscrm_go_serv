package generic

type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id any, entity *T) error
	GetAllWhere(filters map[string]interface{}, entities *[]T) error
	GetAll(entities *[]T) error
	Update(id any, fields map[string]interface{}) error
	Delete(id any, entity *T) error
}
