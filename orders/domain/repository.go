package domain

type Repository interface {
	Create(item any) (any, error)
	Update(item any) error
	FindByID(item any, id uint64) (any, error)
	FindIfODRNotProcessed(orderId uint64) (bool, error)
}
