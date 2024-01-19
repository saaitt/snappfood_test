package domain

type Repository interface {
	Create(item any) (any, error)
	Update(item any) error
	FindByID(item any, id uint64) (any, error)
	FindAllByFilter(item any, key, value any) (any, error)
}
