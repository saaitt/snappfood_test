package kafka_broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func Consumer(server, groupId, offsetReset string, topics []string) *kafka.Consumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupId,
		"auto.offset.reset": offsetReset,
	})

	if err != nil {
		panic(err)
	}

	if err := c.SubscribeTopics(topics, nil); err != nil {
		fmt.Println(err)
		return nil
	}
	return c
}

func Reader(c *kafka.Consumer, ch chan []byte) {
	fmt.Println("starting consumer...")
	run := true
	for run {
		fmt.Println("consumer message loop ...")
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			ch <- msg.Value
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()
}
