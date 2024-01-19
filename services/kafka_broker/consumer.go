package kafka_broker

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"time"
)

func Consumer(server, groupId, offsetReset string, topics []string) *MessageConsumer {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupId,
		"auto.offset.reset": offsetReset,
	})

	if err != nil {
		panic(err)
	}
	fmt.Println(topics)
	if err := c.SubscribeTopics(topics, nil); err != nil {
		fmt.Println(err)
		return nil
	}
	msgCh := make(chan []byte)
	return &MessageConsumer{
		c:     c,
		MsgCh: msgCh,
	}
}

func (b *MessageConsumer) Reader() {
	fmt.Println("starting consumer...")
	run := true
	for run {
		msg, err := b.c.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
			b.MsgCh <- msg.Value
		} else if !err.(kafka.Error).IsTimeout() {
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
	b.c.Close()
}

type MessageConsumer struct {
	MsgCh chan []byte
	c     *kafka.Consumer
}
