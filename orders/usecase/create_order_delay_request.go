package usecase

import (
	"github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/orders/repository"
	"time"
)

func (uc *useCase) CreateOrderDelayReport(req *domain.OrderDelayReportRequest) (*domain.OrderDelayReport, error) {
	var order domain.Order
	res, err := uc.repo.FindByID(&order, req.OrderID)
	if err != nil {
		if err == repository.ErrItemNotFound {
			return nil, domain.ErrorWrongOrderID
		}
		return nil, err
	}
	dt := time.Duration(order.DeliveryTime)
	ExpectedDelivery := order.CreatedAt.UTC().Add(dt * time.Minute)
	if ExpectedDelivery.After(time.Now().UTC()) {
		return nil, domain.ErrorOrderIsOnTheWay
	}
	notProcessed, err := uc.repo.FindIfODRNotProcessed(req.OrderID)
	if err != nil {
		return nil, err
	}
	if notProcessed {
		res, err = uc.repo.Create(&domain.OrderDelayReport{
			OrderID:   req.OrderID,
			CreatedAt: time.Now().UTC(),
		})
		if err != nil {
			return nil, err
		}
		odr := res.(*domain.OrderDelayReport)
		return odr, nil
	}
	return nil, domain.ErrorOrderDelayReportAlreadyExists
}
