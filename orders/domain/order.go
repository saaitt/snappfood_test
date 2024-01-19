package domain

import "time"

type Order struct {
	ID           uint64    `json:"id"`
	DeliveryTime uint64    `json:"delivery_time"`
	CreatedAt    time.Time `json:"created_at"`
	StartAt      time.Time `json:"start_at"`
	VendorID     uint64    `json:"vendor_id"`
}

type OrderRequest struct {
	DeliveryTime uint64 `json:"delivery_time"`
	VendorID     uint64 `json:"vendor_id"`
}

type OrderDelayReport struct {
	ID        uint64    `json:"id"`
	OrderID   uint64    `json:"order_id"`
	Processed bool      `json:"processed"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderDelayReportRequest struct {
	OrderID uint64 `json:"order_id"`
}

type VendorDelayReport struct {
	VendorID    string `json:"vendor_id"`
	DelayCounts uint64 `json:"delay_counts"`
}
