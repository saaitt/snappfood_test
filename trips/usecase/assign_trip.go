package usecase

import (
	"github.com/saaitt/snappfood_test/trips/domain"
	"github.com/saaitt/snappfood_test/trips/repository"
	"time"
)

func (uc *useCase) AssignTrip(req *domain.TripRequest) (*domain.Trip, error) {
	var t domain.Trip
	_, err := uc.repo.FindAllByFilter(&t, "order_id", req.OrderID)
	if err != nil && err != repository.ErrItemNotFound {
		return nil, err
	}
	if !t.Empty() {
		return nil, domain.ErrorOrderIsAlreadyAssigned
	}

	res, err := uc.repo.Create(&domain.Trip{
		OrderID:   req.OrderID,
		Status:    domain.StatusAssigned,
		CreatedAt: time.Now().UTC(),
	})
	if err != nil {
		return nil, err
	}
	trip := res.(*domain.Trip)
	return trip, nil
}
