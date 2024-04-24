package service

type ICrudService[T any] interface {
	Get(id string) (*T, error)
	GetAll() ([]T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id string) error
}
