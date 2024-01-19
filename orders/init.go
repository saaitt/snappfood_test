package orders

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saaitt/snappfood_test/orders/delivery"
	"github.com/saaitt/snappfood_test/orders/domain"
	"github.com/saaitt/snappfood_test/orders/repository"
	"github.com/saaitt/snappfood_test/orders/usecase"
	td "github.com/saaitt/snappfood_test/trips/domain"
	"gorm.io/gorm"
)

func InitUC(repo domain.Repository, tripUc td.UseCase, producer *kafka.Producer, agentTopic string) domain.UseCase {
	return usecase.New(repo, tripUc, producer, agentTopic)
}

func InitRepository(db *gorm.DB) domain.Repository {
	return repository.New(db)
}

func InitHandler(uc domain.UseCase) domain.Handler {
	return delivery.New(uc)
}
