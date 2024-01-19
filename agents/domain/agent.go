package domain

type AgentHistory struct {
	ID            uint64 `json:"id"`
	Processed     bool   `json:"processed"`
	OrderID       uint64 `json:"order_id"`
	DelayReportID uint64 `json:"delay_report_id"`
}
