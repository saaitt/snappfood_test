package main

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/saaitt/snappfood_test/services/kafka_broker"
)

func main() {
	agent := "agent"
	value, _ := json.Marshal("this is an order")
	producer := kafka_broker.Producer("localhost")
	if err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &agent,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil); err != nil {
		fmt.Println(err)
	}
}
