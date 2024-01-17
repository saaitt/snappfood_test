package delivery

import "github.com/saaitt/snappfood_test/orders/domain"

type handler struct {
	uc domain.UseCase
}

func New(uc domain.UseCase) domain.Handler {
	return &handler{
		uc: uc,
	}
}
