package domain

type Order struct {
	ID          string      `json:"id"`
	UserID      string      `json:"userId"`
	Items       []OrderItem `json:"orderItems"`
	TotalAmount float64     `json:"totalAmount"`
}
