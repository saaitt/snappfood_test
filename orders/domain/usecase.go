package domain

type UseCase interface {
	CreateOrder(orderReq *OrderRequest) (*Order, error)
	CreateOrderDelayReport(req *OrderDelayReportRequest) (*OrderDelayReport, error)
}
