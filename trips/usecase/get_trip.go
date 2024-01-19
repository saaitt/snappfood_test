package usecase

import (
	"github.com/saaitt/snappfood_test/trips/domain"
	"github.com/saaitt/snappfood_test/trips/repository"
)

func (uc *useCase) GetTrip(req *domain.TripRequest) (*domain.Trip, error) {
	var t domain.Trip
	_, err := uc.repo.FindAllByFilter(&t, "order_id", req.OrderID)
	if err != nil {
		if err == repository.ErrItemNotFound {
			return nil, domain.TripNotFound
		}
	}
	return &t, nil
}
