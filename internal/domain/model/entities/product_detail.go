package entities

type ProductDetail struct {
	ProductDetailID int64 `json:"product_detail_id"`
	ProductID       int64 `json:"product_id"`
	SizeID          int64 `json:"size_id"`
	ColorID         int64 `json:"color_id"`
	StockCount      int64 `json:"stock_count"`
}
