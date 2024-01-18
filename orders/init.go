package orders

import (
	"github.com/saaitt/snappfood_test/orders/delivery"
	"github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/orders/repository"
	"github.com/saaitt/snappfood_test/orders/usecase"
	td "github.com/saaitt/snappfood_test/trips/domain"
	"gorm.io/gorm"
)

func InitUC(repo domain.Repository, tripUc td.UseCase) domain.UseCase {
	return usecase.New(repo, tripUc)
}

func InitRepository(db *gorm.DB) domain.Repository {
	return repository.New(db)
}

func InitHandler(uc domain.UseCase) domain.Handler {
	return delivery.New(uc)
}
