package service

type ICrudService[T any] interface {
	Get(id int64) (T, error)
	GetAll() ([]T, error)
	Create(entity *T) (T, error)
	Update(entity *T) (T, error)
	Delete(id int64) error
}
