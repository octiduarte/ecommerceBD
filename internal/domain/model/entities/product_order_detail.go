package entities

type ProductOrderDetail struct {
	ProductOrderDetailID int64 `json:"product_order_detail_id"`
	Count                int   `json:"count"`
	ProductID            int64 `json:"product_id"`
	OrderID              int64 `json:"order_id"`
}
