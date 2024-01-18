package app

import (
	orders "github.com/saaitt/snappfood_test/orders"
	trips "github.com/saaitt/snappfood_test/trips"
)

func (a *App) InitOrders() {
	a.orderUc = orders.InitUC(orders.InitRepository(a.db))
	a.orderHr = orders.InitHandler(a.orderUc)
}

func (a *App) InitTrips() {
	uc := trips.InitUC(trips.InitRepository(a.db))
	a.tripHr = trips.InitHandler(uc)
}
