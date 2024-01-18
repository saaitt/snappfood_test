package domain

import "time"

const (
	StatusDelivered = "DELIVERED"
	StatusPicked    = "PICKED"
	StatusAtVendor  = "AT_VENDOR"
	StatusAssigned  = "ASSIGNED"
)

type Trip struct {
	ID        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	OrderID   uint64    `json:"order_id"`
	Status    string    `json:"status"`
}

type TripRequest struct {
	OrderID uint64 `json:"order_id"`
}

type TripUpdateRequest struct {
	ID     uint64 `json:"id"`
	Status string `json:"status"`
}

func (m *Trip) Empty() bool {
	return m.ID < 1
}

func (m *Trip) Ongoing() bool {
	return m.Status == StatusPicked ||
		m.Status == StatusAssigned ||
		m.Status == StatusAtVendor
}
