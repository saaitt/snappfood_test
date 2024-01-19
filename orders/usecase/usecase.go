package usecase

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saaitt/snappfood_test/orders/domain"
	td "github.com/saaitt/snappfood_test/trips/domain"
)

type useCase struct {
	repo       domain.Repository
	tripUc     td.UseCase
	producer   *kafka.Producer
	agentTopic string
}

func New(repo domain.Repository, tuc td.UseCase, producer *kafka.Producer, agentTopic string) domain.UseCase {
	return &useCase{
		repo:       repo,
		tripUc:     tuc,
		producer:   producer,
		agentTopic: agentTopic,
	}
}
