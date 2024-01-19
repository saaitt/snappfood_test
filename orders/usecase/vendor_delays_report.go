package usecase

import "github.com/saaitt/snappfood_test/orders/domain"

func (uc *useCase) VendorDelaysReport() ([]domain.VendorDelayReport, error) {
	return uc.repo.VendorDelaysReport()
}
