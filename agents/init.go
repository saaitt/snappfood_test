package agents

import (
	"github.com/saaitt/snappfood_test/agents/delivery"
	"github.com/saaitt/snappfood_test/agents/domain"
	"github.com/saaitt/snappfood_test/agents/repository"
	"github.com/saaitt/snappfood_test/agents/usecase"
	"github.com/saaitt/snappfood_test/services/kafka_broker"
	"gorm.io/gorm"
)

func InitUC(repo domain.Repository, consumer *kafka_broker.MessageConsumer) domain.UseCase {
	return usecase.New(repo, consumer)
}

func InitRepository(db *gorm.DB) domain.Repository {
	return repository.New(db)
}

func InitHandler(uc domain.UseCase) domain.Handler {
	return delivery.New(uc)
}
