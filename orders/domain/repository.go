package domain

type Repository interface {
	Create(item any) (any, error)
	Update(item any) error
	FindByID(item any, id uint64) error
	FindIfODRNotProcessed(orderId uint64) (bool, error)
	VendorDelaysReport() ([]VendorDelayReport, error)
}
