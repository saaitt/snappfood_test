package orders

import (
	"github.com/saaitt/snappfood_test/orders/delivery"
	"github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/orders/repository"
	"github.com/saaitt/snappfood_test/orders/usecase"
	"gorm.io/gorm"
)

func InitUC(repo domain.Repository) domain.UseCase {
	return usecase.New(repo)
}

func InitRepository(db *gorm.DB) domain.Repository {
	return repository.New(db)
}

func InitHandler(uc domain.UseCase) domain.Handler {
	return delivery.New(uc)
}
