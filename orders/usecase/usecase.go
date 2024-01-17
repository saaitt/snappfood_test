package usecase

import "github.com/saaitt/snappfood_test/orders/domain"

type useCase struct {
	repo domain.Repository
}

func New(repo domain.Repository) domain.UseCase {
	return &useCase{
		repo: repo,
	}
}
