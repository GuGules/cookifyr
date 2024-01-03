package repository

type Repository[T struct{}] interface {
	FindAll() []T
	FindById() T
	Create(T) T
	Delete(T) bool
	Update(T) T
}
