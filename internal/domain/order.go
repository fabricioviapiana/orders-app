package domain

type Order struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Products    []Product `json:"products"`
	TotalAmount float64   `json:"total_amount"`
}
