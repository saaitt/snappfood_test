package domain

import orderDomain "github.com/saaitt/snappfood_test/orders/domain"

type AgentTask struct {
	ID            uint64 `json:"id"`
	Processed     bool   `json:"processed"`
	OrderID       uint64 `json:"order_id"`
	DelayReportID uint64 `json:"delay_report_id"`
}

type AgentTaskResponse struct {
	Task         AgentTask                      `json:"task"`
	Processed    bool                           `json:"processed"`
	Order        orderDomain.Order              `json:"order"`
	DelayReports []orderDomain.OrderDelayReport `json:"delay_reports"`
}

type TaskUpdateRequest struct {
	ID        uint64 `json:"id"`
	Processed bool   `json:"processed"`
}
