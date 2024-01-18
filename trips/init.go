package trips

import (
	"github.com/saaitt/snappfood_test/trips/delivery"
	"github.com/saaitt/snappfood_test/trips/domain"
	"github.com/saaitt/snappfood_test/trips/repository"
	"github.com/saaitt/snappfood_test/trips/usecase"
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
