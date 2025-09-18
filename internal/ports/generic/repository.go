package generic

type Repository[T any] interface {
	Create(entity *T) error
	GetByID(id any, entity *T) error
	GetAllWhere(fields []string, values []string, entities *[]T) error
	GetAll(entities *[]T) error
	Update(entity *T) error
	Delete(entity *T) error
}
