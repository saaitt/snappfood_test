package usecase

import (
	"github.com/saaitt/snappfood_test/agents/domain"
	"github.com/saaitt/snappfood_test/services/kafka_broker"
)

type useCase struct {
	repo     domain.Repository
	consumer *kafka_broker.MessageConsumer
}

func New(repo domain.Repository, consumer *kafka_broker.MessageConsumer) domain.UseCase {
	return &useCase{
		repo:     repo,
		consumer: consumer,
	}
}
