package entities

import "time"

type ProductOrder struct {
	OrderID   int64     `json:"order_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    int64     `json:"user_id"`
}
