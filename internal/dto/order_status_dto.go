package dto

type OrderStatus string

const (
	OrderStatusCreated  OrderStatus = "created"
	OrderStatusSent     OrderStatus = "sent"
	OrderStatusAccepted OrderStatus = "accepted"
	OrderStatusRejected OrderStatus = "rejected"
)

var validOrderStatuses = map[OrderStatus]bool{
	OrderStatusCreated:  true,
	OrderStatusSent:     true,
	OrderStatusAccepted: true,
	OrderStatusRejected: true,
}

func (s OrderStatus) String() string {
	return string(s)
}

// IsValid checks if the OrderStatus is valid
func (s OrderStatus) IsValid() bool {
	_, exists := validOrderStatuses[s]
	return exists
}
