package usecase

import (
	"github.com/saaitt/snappfood_test/orders/domain"
	"time"
)

func (uc *useCase) CreateOrder(orderReq *domain.OrderRequest) (*domain.Order, error) {
	res, err := uc.repo.Create(&domain.Order{
		DeliveryTime: orderReq.DeliveryTime,
		CreatedAt:    time.Now().UTC(),
		StartAt:      time.Now().UTC(),
		VendorID:     orderReq.VendorID,
	})
	if err != nil {
		return nil, err
	}
	order := res.(*domain.Order)
	return order, nil
}
