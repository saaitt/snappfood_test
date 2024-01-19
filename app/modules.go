package app

import (
	agents "github.com/saaitt/snappfood_test/agents"
	orders "github.com/saaitt/snappfood_test/orders"
	trips "github.com/saaitt/snappfood_test/trips"
)

func (a *App) InitOrders() {
	a.orderUc = orders.InitUC(
		orders.InitRepository(a.db),
		a.tripUc,
		a.producer,
		a.config.Broker.AgentTopic,
	)
	a.orderHr = orders.InitHandler(a.orderUc)
}

func (a *App) InitTrips() {
	a.tripUc = trips.InitUC(trips.InitRepository(a.db))
	a.tripHr = trips.InitHandler(a.tripUc)
}

func (a *App) InitAgents() {
	a.agentUc = agents.InitUC(agents.InitRepository(a.db), a.consumer)
	a.agentHr = agents.InitHandler(a.agentUc)
}
