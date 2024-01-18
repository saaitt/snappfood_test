package usecase

import (
	"github.com/saaitt/snappfood_test/orders/domain"
	td "github.com/saaitt/snappfood_test/trips/domain"
)

type useCase struct {
	repo   domain.Repository
	tripUc td.UseCase
}

func New(repo domain.Repository, tuc td.UseCase) domain.UseCase {
	return &useCase{
		repo:   repo,
		tripUc: tuc,
	}
}
