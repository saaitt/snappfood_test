package usecase

import (
	"github.com/saaitt/snappfood_test/trips/domain"
	"github.com/saaitt/snappfood_test/trips/repository"
)

func (uc *useCase) UpdateTrip(req *domain.TripUpdateRequest) (*domain.Trip, error) {
	var t domain.Trip
	_, err := uc.repo.FindByFilter(&t, "id", req.ID)
	if err != nil {
		if err == repository.ErrItemNotFound {
			return nil, domain.TripNotFound
		}
		return nil, err
	}

	switch t.Status {
	case domain.StatusAssigned:
		if req.Status != domain.StatusAtVendor {
			return nil, domain.ErrorTripStatusCanOnlyChangeToAtVendor
		}
	case domain.StatusAtVendor:
		if req.Status != domain.StatusPicked {
			return nil, domain.ErrorTripStatusCanOnlyChangeToPicked
		}
	case domain.StatusPicked:
		if req.Status != domain.StatusDelivered {
			return nil, domain.ErrorTripStatusCanOnlyChangeToDelivered
		}
	default:
	}

	t.Status = req.Status
	return &t, uc.repo.Update(&t)
}
