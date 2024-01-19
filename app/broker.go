package app

import "github.com/saaitt/snappfood_test/services/kafka_broker"

func (a *App) InitBroker() {
	a.consumer = kafka_broker.Consumer(
		a.config.Broker.Server,
		a.config.Broker.GroupId,
		a.config.Broker.OffsetReset,
		a.config.Broker.Topics)
	a.producer = kafka_broker.Producer(a.config.Broker.Server)
	go a.consumer.Reader()

}
