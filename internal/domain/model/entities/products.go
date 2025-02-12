package entities

type Product struct {
	ProductID       int64   `json:"product_id"`
	Name            string  `json:"name"`
	MainDescription string  `json:"main_description"`
	LongDescription string  `json:"long_description"`
	Price           float64 `json:"price"`
	CategoryID      int64   `json:"category_id"`
	StoreID         int64   `json:"store_id"`
	DiscountID      int64   `json:"discount_id"`
}
