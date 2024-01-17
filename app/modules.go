package app

import (
	o "github.com/saaitt/snappfood_test/orders"
)

func (a *App) InitOrders() {
	a.orderUc = o.InitUC(o.InitRepository(a.db))
	a.orderHr = o.InitHandler(a.orderUc)
}
