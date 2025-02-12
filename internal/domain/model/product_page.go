package model

type ProductPage struct {
	StoreName       string   `json:"store_name"`
	ProductID       int64    `json:"product_id"`
	Name            string   `json:"name"`
	MainDescription string   `json:"main_description"`
	LongDescription string   `json:"long_description"`
	Price           float64  `json:"price"`
	CategoryName    string   `json:"category_name"`
	DiscountType    *string  `json:"discount_type"`
	DiscountAmount  *int64   `json:"discount_amount"`
	StockCount      int64    `json:"stock_count"`
	Sizes           []string `json:"sizes"`
	Colors          []string `json:"colors"`
	Images          []string `json:"images"`
}
